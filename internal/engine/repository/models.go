package repository

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name string `gorm:"unique"`
	Type string
	Args datatypes.JSON
}

type BackupPolicy struct {
	gorm.Model
	ResourceID uint
	Retention  int
	Status     string
}

type BackupSet struct {
	gorm.Model
	IsValid    bool `gorm:"default:false"`
	Size       int64
	BackupTime time.Time
	ResourceID uint
	Retention  int
}

type ScheduledJob struct {
	gorm.Model

	JobType      string
	Status       string
	ResourceID   uint
	ResourceType string
	Args         datatypes.JSON

	BackupSetID uint
}
