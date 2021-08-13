package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"NintendoCenter/nintendo-parser/internal/protos"
	"NintendoCenter/nintendo-parser/internal/provider"
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
	err = container.Invoke(func(s *service.WorkScheduler, l *zap.Logger) {
		go func() {
			s.Start(ctx, func(game *protos.Game) {
				l.Info(fmt.Sprintf("got %s (%s) game as parse result", game.Title, game.Id))
			})
		}()
	})
	if err != nil {
		t1.Fatal("error caught while doing parsing task", zap.Error(err))
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals
	t1.Info("caught system signal. Terminating...")
}
