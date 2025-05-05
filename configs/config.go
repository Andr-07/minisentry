package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db    DbConfig
	Kafka KafkaConfig
}

type DbConfig struct {
	Dsn string
}

type KafkaConfig struct {
	Broker string
	Topic string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Kafka: KafkaConfig{
			Broker: os.Getenv("KAFKA_BROKER"),
			Topic: "logs",
		},
	}
}
