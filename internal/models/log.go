package models

import (
	"gorm.io/datatypes"
	"time"
)

type LogEntry struct {
	ID      uint           `gorm:"primaryKey"`
	Message string
	Level   string
	Meta    datatypes.JSON
	Time    time.Time
}