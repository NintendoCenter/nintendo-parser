package service

import "go.uber.org/zap"

type offsetParams struct {
	offset int
	limit  int
}

type WorkScheduler struct {
	itemLimit          int
	logger             *zap.Logger
	maxWorkers         int
	requestInterval    int
	parseIntervalHours int
}

func NewWorkScheduler(
	itemLimit int,
	maxWorkers int,
	requestInterval int,
	parseIntervalHours int,
	logger *zap.Logger,
) *WorkScheduler {
	return &WorkScheduler{
		itemLimit:          itemLimit,
		logger:             logger,
		maxWorkers:         maxWorkers,
		requestInterval:    requestInterval,
		parseIntervalHours: parseIntervalHours,
	}
}
