package service

import "go.uber.org/zap"

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
}

func NewWorkScheduler(
	itemLimit int,
	maxWorkers int,
	requestInterval int,
	parseIntervalHours int,
	parser *Parser,
	logger *zap.Logger,
) *WorkScheduler {
	return &WorkScheduler{
		itemLimit:          itemLimit,
		maxWorkers:         maxWorkers,
		requestInterval:    requestInterval,
		parseIntervalHours: parseIntervalHours,
		parser:             parser,
		logger:             logger,
	}
}
