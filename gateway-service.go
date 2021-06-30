package main

import (
	"flag"
	"os"
	"github.com/Kowiste/kafkasample/service"

)

func main() {
	g := service.New()
	c := ReadConfig()
    g.Start(*c)
}

func ReadConfig() *service.Config {
	conf := new(service.Config)
	//ENVIOREMEN VARIABLE
	KafkaAddrEnv := os.Getenv("KAFKA_ADDR")
	KafkaTopicEnv := os.Getenv("KAFKA_TOPIC")

	//COMMAND LINE ARGUMENT
	KafkaAddr := flag.String("ka", "localhost", "Address of Kafka")
	KafkaTopic := flag.String("kt", "measure", "Kafka Topic")

	flag.Parse()
	if KafkaAddrEnv == "" {
		conf.Kafka.Address = *KafkaAddr
	} else {
		conf.Kafka.Address = KafkaAddrEnv
	}
	if KafkaTopicEnv == "" {
		conf.Kafka.Topic = *KafkaTopic
	} else {
		conf.Kafka.Topic = KafkaTopicEnv
	}
	return conf
}

