package syncbyte

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/skyline93/ctask"
	"gorm.io/gorm"
)

const (
	TypeBackup = "ctask:backup"
)

type BackupJobPayload struct {
	ScheduledJobID uint
	ResourceID     uint
	Retention      int
}

func NewBackupTask(ScheduledJobID uint, resourceID uint, retention int) (*ctask.Task, error) {
	payload, err := json.Marshal(BackupJobPayload{ScheduledJobID: ScheduledJobID, ResourceID: resourceID, Retention: retention})
	if err != nil {
		return nil, err
	}
	return ctask.NewTask(TypeBackup, payload), nil
}

type BackupHandler struct {
	db *gorm.DB
}

func NewBackupHandler(db *gorm.DB) *BackupHandler {
	return &BackupHandler{db: db}
}

func (h *BackupHandler) ProcessTask(ctx context.Context, t *ctask.Task) error {
	var p BackupJobPayload

	log.Println("start unmarshal payload")
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	log.Printf("process task %v", p)
	sj, err := GetScheduledJob(h.db, int(p.ScheduledJobID))
	if err != nil {
		log.Printf("get scheduled job failed, err: %v", err)
		return err
	}

	j := BackupJob{ScheduledJob: *sj}
	if err := j.Run(h.db, p.ResourceID, p.Retention); err != nil {
		log.Printf("running job failed, err: %v", err)
		return err
	}

	return nil
}
