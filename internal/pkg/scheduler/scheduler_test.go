package scheduler_test

import (
	"context"
	"testing"
	"time"

	"github.com/gorhill/cronexpr"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/pkg/scheduler"
)

type TJob struct {
	Cron         string
	nextFireTime *time.Time
}

func (j *TJob) Execute() error {
	log.Debugf("execute job, %s", j.Key())
	return nil
}

func (j *TJob) Key() string {
	return uuid.NewV4().String()
}

func (j *TJob) ShouldStart() bool {
	if j.nextFireTime == nil {
		j.SetNextFireTime()
	}
	return time.Now().After(*j.nextFireTime)
}

func (j *TJob) SetNextFireTime() {
	if j.nextFireTime == nil {
		now := time.Now()
		j.nextFireTime = &now
		return
	}

	nt := cronexpr.MustParse(j.Cron).Next(time.Now())
	j.nextFireTime = &nt
}

func TestJob(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	ctx := context.TODO()
	sch := scheduler.New(ctx, 1)

	go sch.Start()

	// sch.AddJob(&TJob{})
	sch.AddPeriodicalJob(&TJob{Cron: "* * * * *"})

	time.Sleep(time.Second * 60 * 2)
}
