package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/grpc"
	"github.com/skyline93/syncbyte-go/internal/engine/schema"
	"github.com/skyline93/syncbyte-go/pkg/mongodb"
)

type BackupJob struct {
	ScheduledJob

	BackupSetID uint
}

func (b *BackupJob) execute() (err error) {
	backupTime := time.Now()

	backupset, err := backup.GetBackupSet(b.BackupSetID)
	if err != nil {
		return err
	}

	log.Infof("start backup job, job_id: %d", b.ID)
	if err = b.start(); err != nil {
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
	if b.ResourceType == "nas" {
		args := backup.NasResourceArgs{}
		if err = json.Unmarshal(b.Args, &args); err != nil {
			return err
		}

		backup_path = args.Dir
	}

	log.Debug("send backup job to grpc server")
	go grpc.Backup(fiChan, backup_path, ctx)

	col := mongoClient.GetCollection(fmt.Sprintf("backupset-%d", b.BackupSetID))

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
	if err = b.complete(); err != nil {
		return err
	}

	return nil
}

func (b *BackupJob) Execute() error {
	return b.execute()
}
