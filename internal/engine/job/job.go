package job

import (
	"time"

	"github.com/gorhill/cronexpr"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/pkg/scheduler"
)

type BaseJob struct {
	nextFireTime *time.Time
	Scheduler    *scheduler.Scheduler
}

func (j *BaseJob) Execute() error {
	log.Warn("you must overwrite the function!!!")
	return nil
}

func (j *BaseJob) Key() string {
	return uuid.NewV4().String()
}

type BasePeriodicalJob struct {
	BaseJob
	Cron string
}

func (pj *BasePeriodicalJob) ShouldStart() bool {
	if pj.nextFireTime == nil {
		pj.SetNextFireTime()
	}
	return time.Now().After(*pj.nextFireTime)
}

func (pj *BasePeriodicalJob) SetNextFireTime() {
	if pj.nextFireTime == nil {
		now := time.Now()
		pj.nextFireTime = &now
		return
	}

	nt := cronexpr.MustParse(pj.Cron).Next(time.Now())
	pj.nextFireTime = &nt
}

type ScheduledJob struct {
	BaseJob
	ID           uint
	JobType      string
	ResourceID   uint
	ResourceType string
	Args         []byte
}

func (s *ScheduledJob) start() error {
	if result := repository.Repo.Model(&repository.ScheduledJob{}).Where("id = ?", s.ID).Update("status", "running"); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ScheduledJob) complete() error {
	if result := repository.Repo.Model(&repository.ScheduledJob{}).Where("id = ?", s.ID).Update("status", "completed"); result.Error != nil {
		return result.Error
	}
	return nil
}