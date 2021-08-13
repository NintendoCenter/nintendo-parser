package provider

import (
	"NintendoCenter/nintendo-parser/config"
	"NintendoCenter/nintendo-parser/internal/client"
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
		func() *client.Nintendo {
			return client.NewNintendoClient()
		},
		func(c *client.Nintendo) *service.Parser {
			return service.NewParser(c)
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
		config.NewConfig,
	}

	for _, i := range constructors {
		if err := container.Provide(i); err != nil {
			return nil, err
		}
	}

	return container, nil
}
