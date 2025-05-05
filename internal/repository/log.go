package repository

import (
	"minisentry/internal/db"
	"minisentry/internal/models"
)

type LogRepository struct {
	Db *db.Db
}

func NewLogRepository(db *db.Db) *LogRepository {
	return &LogRepository{
		Db: db,
	}
}

func (repo *LogRepository) Save(log *models.LogEntry) (*models.LogEntry, error) {
	result := repo.Db.DB.Create(log)
	if result.Error != nil {
		return nil, result.Error
	}
	return log, nil
}