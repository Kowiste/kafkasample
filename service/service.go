package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Kowiste/kafkasample/handler/kafka"
	log "github.com/sirupsen/logrus"
)

type GatewayServer struct {
	Kafka *kafka.Kafka
}

type Config struct {
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Address  string
	Topic    string
	User     string
	Password string
}

func New() *GatewayServer {
	return &GatewayServer{}
}

func (g *GatewayServer) Start(config Config) error {
	g.Kafka = kafka.New()
	conKafka := kafka.Config{
		Endpoint: config.Kafka.Address,
		Topic:    config.Kafka.Topic,
	}
	ctx := context.Background()
	g.Kafka.Start(conKafka)
	go g.Kafka.Consume(ctx)
	ticker := time.NewTicker(2000 * time.Millisecond)
	payload := struct {
		Name  string
		Value int
	}{
		Name:  "medida 1",
		Value: 0,
	}

	for {
		select {
		case <-ticker.C:
			payload.Value++
			log.Infoln("Publishing ... ", payload.Value)
			
			bytes, _ := json.Marshal(payload)
			err := g.Kafka.Write([]byte("measure"), bytes)
			if err != nil {
				log.Errorln(err)
			}
			if payload.Value > 300 {
				payload.Value = 0
			}
		}
	}
}
