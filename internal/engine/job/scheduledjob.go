package job

import (
	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/pkg/scheduler"
)

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

type Scheduler struct {
	BasePeriodicalJob
}

func NewJobScheduler(sch *scheduler.Scheduler, cron string) *Scheduler {
	js := &Scheduler{}
	js.Scheduler = sch
	js.Cron = cron

	return js
}

func (jsch *Scheduler) Execute() error {
	return jsch.run()
}

func (jsch *Scheduler) getTodoScheduledJobs() (scheduledJobs []repository.ScheduledJob, err error) {
	var items []repository.ScheduledJob

	if result := repository.Repo.Where("status = ?", "queued").Find(&items); result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (jsch *Scheduler) run() error {
	jobs, err := jsch.getTodoScheduledJobs()
	if err != nil {
		return err
	}

	for _, j := range jobs {
		var scheduledJob scheduler.Job

		switch j.JobType {
		case "backup":
			scheduledJob = &BackupJob{
				ScheduledJob: ScheduledJob{
					ID:           j.ID,
					JobType:      j.JobType,
					ResourceID:   j.ResourceID,
					ResourceType: j.ResourceType,
					Args:         j.Args,
				},
				BackupSetID: j.BackupSetID,
			}
		}

		log.Infof("job scheduler execute job, id: %v", j.ID)
		go jsch.Scheduler.AddJob(scheduledJob)
		log.Infof("try to child job")
	}

	return nil
}
