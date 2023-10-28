package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/skyline93/ctask"
)

const (
	TypeBackup = "ctask:backup"
)

type BackupJobPayload struct {
	BackupSetID uint
}

func NewBackupTask(backupSetID uint) (*ctask.Task, error) {
	payload, err := json.Marshal(BackupJobPayload{BackupSetID: backupSetID})
	if err != nil {
		return nil, err
	}
	return ctask.NewTask(TypeBackup, payload), nil
}

func HandleBackupTask(ctx context.Context, t *ctask.Task) error {
	var p BackupJobPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("parse task payload failed: %v", err)
	}
	log.Printf("handle backup task: backup_set_id=%d", p.BackupSetID)

	// TODO

	return nil
}
