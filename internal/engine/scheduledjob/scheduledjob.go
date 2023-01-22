package scheduledjob

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/grpc"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/engine/schema"
	"github.com/skyline93/syncbyte-go/pkg/mongodb"
)

type ScheduledJob struct {
	ID           uint
	JobType      string
	ResourceID   uint
	ResourceType string
	Args         []byte

	BackupSetID uint
}

func (s *ScheduledJob) Run() error {
	switch s.JobType {
	case "backup":
		log.Infof("scheduler start backup job")
		if err := s.runBackupJob(); err != nil {
			return err
		}
	}

	return nil
}

func (s *ScheduledJob) runBackupJob() (err error) {
	backupTime := time.Now()

	backupset, err := backup.GetBackupSet(s.BackupSetID)
	if err != nil {
		return err
	}

	log.Infof("start backup job, job_id: %d", s.ID)
	if err = s.start(); err != nil {
		return err
	}

	log.Infof("set backup time %s", backupTime)
	if err = backupset.SetBackupTime(backupTime); err != nil {
		return err
	}

	mongoClient, err := mongodb.NewClient(config.Conf.Core.MongodbUri)
	if err != nil {
		return err
	}
	defer mongoClient.Close()

	fiChan := make(chan schema.FileInfo)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var backup_path string
	if s.ResourceType == "nas" {
		args := backup.NasResourceArgs{}
		if err = json.Unmarshal(s.Args, &args); err != nil {
			return err
		}

		backup_path = args.Dir
	}

	log.Debug("send backup job to grpc server")
	go grpc.Backup(fiChan, backup_path, ctx)

	col := mongoClient.GetCollection(fmt.Sprintf("backupset-%d", s.BackupSetID))

	for fi := range fiChan {
		if _, err := col.InsertOne(context.TODO(), fi); err != nil {
			log.Errorf("insert error, err: %v", err)
			continue
		}

		log.Debugf("fi: %s", fi.String())
	}

	log.Infof("set backupset to available")
	if err = backupset.SetAvailability(true); err != nil {
		return err
	}

	log.Infof("complete backup job")
	if err = s.complete(); err != nil {
		return err
	}

	return nil
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

func GetTodoScheduledJobs() (scheduledJobs []ScheduledJob, err error) {
	var items []repository.ScheduledJob

	if result := repository.Repo.Where("status = ?", "queued").Find(&items); result.Error != nil {
		return nil, result.Error
	}

	for _, j := range items {
		sj := ScheduledJob{
			ID:           j.ID,
			JobType:      j.JobType,
			ResourceID:   j.ResourceID,
			ResourceType: j.ResourceType,
			Args:         j.Args,
			BackupSetID:  j.BackupSetID,
		}
		scheduledJobs = append(scheduledJobs, sj)
	}

	return scheduledJobs, nil
}
