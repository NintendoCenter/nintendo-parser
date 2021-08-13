package provider

import (
	"NintendoCenter/nintendo-parser/config"
	"NintendoCenter/nintendo-parser/internal/client"
	"NintendoCenter/nintendo-parser/internal/queue/producer"
	"NintendoCenter/nintendo-parser/internal/service"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	constructors := []interface{}{
		func(cfg *config.Config) (*zap.Logger, error) {
			return NewLogger(cfg.LogLevel)
		},
		func(l *zap.Logger) *client.Nintendo {
			return client.NewNintendoClient(l)
		},
		func(c *client.Nintendo, l *zap.Logger) *service.Parser {
			return service.NewParser(c, l)
		},
		func(cfg *config.Config, l *zap.Logger, p *service.Parser) *service.WorkScheduler {
			return service.NewWorkScheduler(
				cfg.ItemLimit,
				cfg.MaxWorkers,
				cfg.RequestInterval,
				cfg.ParseIntervalHours,
				p,
				l,
			)
		},
		func(cfg *config.Config, l *zap.Logger) (*producer.GameProducer, error) {
			return producer.NewProducer(cfg.QueueAddr, cfg.GamesTopic, l)
		},
		config.NewConfig,
	}

	for _, i := range constructors {
		if err := container.Provide(i); err != nil {
			return nil, err
		}
	}

	return container, nil
}
