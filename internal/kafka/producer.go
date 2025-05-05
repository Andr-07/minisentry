package kafka

import (
	"context"
	"log"
	"minisentry/configs"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(conf *configs.KafkaConfig) *KafkaProducer {
	return &KafkaProducer{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(conf.Broker),
			Topic:    conf.Topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (k *KafkaProducer) WriteMessage(data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := k.Writer.WriteMessages(ctx, kafka.Message{Value: data})
	if err != nil {
		log.Println("Kafka write error:", err)
	} else {
		log.Println("Message sent to Kafka:", string(data))
	}
	return err
}
