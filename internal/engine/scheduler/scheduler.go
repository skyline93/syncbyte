package scheduler

import (
	"log"
	"time"

	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/scheduling"
	"gorm.io/gorm"
)

const (
	TypeCron      = "cron"
	TypeFrequency = "frequency"
)

type BackupSchedule struct {
	gorm.Model
	PolicyID  uint
	Type      string
	Frequency int
	Cron      string
	NextTime  *time.Time
	IsActive  bool
}

func (bs *BackupSchedule) IsReached() bool {
	if bs.NextTime == nil {
		return true
	}

	return bs.NextTime.Before(time.Now())
}

func LoadBackupSchedules(db *gorm.DB) ([]BackupSchedule, error) {
	var schedules []BackupSchedule
	if result := db.Where("is_active = ?", true).Find(&schedules); result.Error != nil {
		return nil, result.Error
	}

	return schedules, nil
}

type Scheduler struct {
	db *gorm.DB
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Run() error {
	schs, err := LoadBackupSchedules(s.db)
	if err != nil {
		return err
	}

	for _, sch := range schs {
		if !sch.IsReached() {
			continue
		}

		pl := backup.GetPolicy(sch.PolicyID)
		jobID, err := scheduling.ScheduleBackup(pl.ResourceID)
		if err != nil {
			log.Printf("schedule backup job failed, error: %v", err)
			continue
		}

		log.Printf("schedule one backup job %d", jobID)
	}

	return nil
}
