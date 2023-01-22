package scheduler

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/pkg/worker"
)

type Job interface {
	Execute() error
	Key() string
}

type Trigger interface {
	ShouldStart() bool
	SetNextFireTime()
}

type PeriodicalJob interface {
	Job
	Trigger
}

type Scheduler struct {
	ctx                 context.Context
	wp                  *worker.Pool
	maxWorkerConcurrent int

	periodicalJobs sync.Map
}

func New(ctx context.Context, maxWorkerConcurrent int) *Scheduler {
	return &Scheduler{
		ctx:                 ctx,
		wp:                  worker.NewPool(ctx, maxWorkerConcurrent),
		maxWorkerConcurrent: maxWorkerConcurrent,
	}
}

func (s *Scheduler) AddJob(job Job) {
	s.wp.Submit(job)
}

func (s *Scheduler) AddPeriodicalJob(job PeriodicalJob) {
	log.Debugf("add periodical job to scheduler")
	s.periodicalJobs.Store(job.Key(), job)
}

func (s *Scheduler) Start() {
	for {
		select {
		case <-time.NewTicker(time.Second * 10).C:
			log.Debugf("ticker run... periodical jobs")
			s.periodicalJobs.Range(func(k, v any) bool {
				j, ok := v.(PeriodicalJob)
				if !ok {
					return false
				}

				if j.ShouldStart() {
					s.AddJob(j)
					j.SetNextFireTime()
				}

				return true
			})
		case <-s.ctx.Done():
			return
		}
	}
}
