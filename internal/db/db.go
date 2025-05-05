package db

import (
	"log"
	"minisentry/configs"
	"minisentry/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.DbConfig) *Db {
	db, err := gorm.Open(postgres.Open(conf.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err := db.AutoMigrate(&models.LogEntry{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return &Db{db}
}