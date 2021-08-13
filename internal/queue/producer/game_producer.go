package producer

import (
	"encoding/json"

	"NintendoCenter/nintendo-parser/internal/protos"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

type GameProducer struct {
	topic string
	producer *nsq.Producer
	logger *zap.Logger
}

func NewProducer(queueAddress string, topic string, logger *zap.Logger) (*GameProducer, error) {
	producer, err := nsq.NewProducer(queueAddress, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	return &GameProducer{
		topic: topic,
		producer: producer,
		logger: logger,
	}, nil
}

func (p *GameProducer) SendGames(games []*protos.Game) error {
	message := make([][]byte, 0, len(games))
	for _, game := range games {
		g, err := json.Marshal(game)
		if err != nil {
			p.logger.Fatal("cannot produce message with a game data")
			return err
		}
		message = append(message, g)
	}

	return p.producer.MultiPublish(p.topic, message)
}

func (p *GameProducer) Stop() {
	p.logger.Info("producer stopped manually")
	p.producer.Stop()
}