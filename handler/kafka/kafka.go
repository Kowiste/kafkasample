package kafka

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Kafka struct {
	Context context.Context
	Writer  *kafka.Writer
	Config  Config
}
type Config struct {
	Endpoint string
	Topic    string
}

func New() *Kafka {
	return &Kafka{}
}

func (k *Kafka) Start(config Config) error {
	k.Config = config
	k.Context = context.Background()
	k.Writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{k.Config.Endpoint},
		Topic:   k.Config.Topic,
	})
	return nil

}
func (k *Kafka) Write(key, Data []byte) error {
	errs := k.Writer.WriteMessages(k.Context, kafka.Message{
		Key:   key,
		Value: Data,
	})
	return errs
}
func (k *Kafka) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{k.Config.Endpoint}, // broker2Address, broker3Address},
		Topic:   k.Config.Topic,
		GroupID: "my-group3",
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("could not read message " + err.Error())
		}
		log.Infoln("received: ", string(msg.Value))
	}
}
