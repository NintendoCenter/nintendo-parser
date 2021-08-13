package service

import (
	"context"
	"fmt"
	"time"

	"NintendoCenter/nintendo-parser/internal/protos"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type offsetParams struct {
	offset int
	limit  int
}

type WorkScheduler struct {
	itemLimit          int
	maxWorkers         int
	requestInterval    int
	parseIntervalHours int
	parser             *Parser
	logger             *zap.Logger
	cron               *gocron.Scheduler
}

func NewWorkScheduler(
	itemLimit int,
	maxWorkers int,
	requestInterval int,
	parseIntervalHours int,
	parser *Parser,
	logger *zap.Logger,
) *WorkScheduler {
	c := gocron.NewScheduler(time.UTC)
	c.TagsUnique()
	return &WorkScheduler{
		itemLimit:          itemLimit,
		maxWorkers:         maxWorkers,
		requestInterval:    requestInterval,
		parseIntervalHours: parseIntervalHours,
		parser:             parser,
		logger:             logger,
		cron:               c,
	}
}

const JobTag = "parser"

func (s *WorkScheduler) Start(ctx context.Context, onGameLoaded func(games []*protos.Game)) {
	cmdChan := make(chan offsetParams, s.maxWorkers)
	gr, ctx := errgroup.WithContext(ctx)
	for i := 0; i < s.maxWorkers; i++ {
		gr.Go(func() error {
			for true {
				select {
				case <- ctx.Done():
					s.logger.Info("worker stopped by the context")
					return nil
				case params := <- cmdChan:
					gameList, err := s.parser.ParseGames(params.offset, params.limit)
					if err != nil {
						return err
					}
					if len(gameList) == 0 {
						s.logger.Info(fmt.Sprintf("empty response from parser. stopping job at offset %d", params.offset))
						_ = s.cron.RemoveByTag(JobTag)
						continue
					}
					s.logger.Debug(fmt.Sprintf("Loaded %d items with offset %d and limit %d", len(gameList), params.offset, params.limit))
					onGameLoaded(gameList)
				}
			}
			return nil
		})
	}

	s.cron.Every(s.parseIntervalHours).Hours().Do(func() {
		params := offsetParams{
			offset: 0,
			limit:  s.itemLimit,
		}
		_ = s.cron.RemoveByTag(JobTag)
		s.cron.Every(s.requestInterval).Seconds().Tag(JobTag).Do(func() {
			cmdChan <- params
			params.offset += params.limit
		})
	})

	s.cron.StartAsync()

	if err := gr.Wait(); err != nil {
		s.logger.Fatal(err.Error())
	}

	s.cron.Stop()
}
