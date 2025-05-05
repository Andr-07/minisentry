package kafka

import (
	"context"
	"encoding/json"
	"log"
	"minisentry/configs"
	"minisentry/internal/models"
	"minisentry/internal/repository"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader        *kafka.Reader
	LogRepository *repository.LogRepository
}

func NewKafkaConsumer(conf *configs.KafkaConfig, logRepository *repository.LogRepository) *KafkaConsumer {
	return &KafkaConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{conf.Broker},
			Topic:   conf.Topic,
			GroupID: uuid.NewString(),
		}),
		LogRepository: logRepository,
	}
}

func (k *KafkaConsumer) ReadAll() error {
	for {
		log.Println("Waiting for message from Kafka...")

		msg, err := k.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Error reading message:", err)
			continue
		}

		log.Printf("✅ Received message from Kafka: topic=%s partition=%d offset=%d value=%s",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value),
		)

		var logEntry models.LogEntry
		if err := json.Unmarshal(msg.Value, &logEntry); err != nil {
			log.Println("❌ JSON unmarshal error:", err)
			continue
		}

		log.Printf("Parsed LogEntry: %+v", logEntry)
		logEntry.Time = time.Now()

		if _, err := k.LogRepository.Save(&logEntry); err != nil {
			log.Println("❌ Failed to save to DB:", err)
		} else {
			log.Println("✅ Saved log entry to DB")
		}
	}
}
