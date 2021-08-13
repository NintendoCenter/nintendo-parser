package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"NintendoCenter/nintendo-parser/internal/protos"
	"NintendoCenter/nintendo-parser/internal/provider"
	"NintendoCenter/nintendo-parser/internal/queue/producer"
	"NintendoCenter/nintendo-parser/internal/service"
	"go.uber.org/zap"
)

func main() {
	t1, _ := zap.NewProduction()
	container, err := provider.BuildContainer()
	if err != nil {
		t1.Fatal("cannot start service due to dependencies error", zap.Error(err))
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = container.Invoke(func(s *service.WorkScheduler, p *producer.GameProducer, l *zap.Logger) {
		go func() {
			s.Start(ctx, func(games []*protos.Game) {
				if err := p.SendGames(games); err != nil {
					l.Fatal("error on sending games to the queue", zap.Error(err))
				}
			})
		}()
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
		defer signal.Stop(signals)
		defer p.Stop()

		<-signals
		t1.Info("caught system signal. Terminating...")
	})
	if err != nil {
		t1.Fatal("error caught while doing parsing task", zap.Error(err))
	}
}
