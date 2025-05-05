package main

import (
	"log"
	"minisentry/configs"
	"minisentry/internal/db"
	"minisentry/internal/kafka"
	"minisentry/internal/repository"
)

func main() {
	conf := configs.LoadConfig()
	log.Println("Starting Kafka consumer...")

	db := db.NewDb(&conf.Db)
	logRepository := repository.NewLogRepository(db)
	k := kafka.NewKafkaConsumer(&conf.Kafka, logRepository)
	
	defer k.Reader.Close()
	k.ReadAll()
}
