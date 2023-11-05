package syncbyte

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JobType string

const (
	JobTypeBackup JobType = "backup"
)

type JobStatus string

const (
	JobStatusQueued    JobStatus = "queued"
	JobStatusRunning   JobStatus = "running"
	JobStatusFailed    JobStatus = "failed"
	JobStatusSuccessed JobStatus = "successed"
)

type ScheduledJob struct {
	gorm.Model
	JobType        JobType
	Status         JobStatus
	StartTime      *time.Time
	EndTime        *time.Time
	BackupPolicyID uint
	ResourceAttr   datatypes.JSONMap
}

func CreateBackupJob(db *gorm.DB, backupPolicy *BackupPolicy) (*ScheduledJob, error) {
	j := ScheduledJob{JobType: JobTypeBackup, BackupPolicyID: backupPolicy.ID, ResourceAttr: backupPolicy.Resource.Attributes, Status: JobStatusQueued}
	if result := db.Create(&j); result.Error != nil {
		return nil, result.Error
	}

	return &j, nil
}

func GetScheduledJob(db *gorm.DB, jobID int) (*ScheduledJob, error) {
	var j ScheduledJob
	if result := db.Where("id = ?", jobID).First(&j); result.Error != nil {
		return nil, result.Error
	}

	return &j, nil
}

func (j *ScheduledJob) Start(db *gorm.DB) error {
	if result := db.Where("id = ?", j.ID).Updates(map[string]interface{}{"status": JobStatusRunning, "start_time": time.Now()}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (j *ScheduledJob) Fail(db *gorm.DB, error_message string) error {
	if result := db.Where("id = ?", j.ID).Updates(map[string]interface{}{"status": JobStatusFailed, "end_time": time.Now(), "error_message": error_message}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (j *ScheduledJob) Success(db *gorm.DB) error {
	if result := db.Where("id = ?", j.ID).Updates(map[string]interface{}{"status": JobStatusSuccessed, "end_time": time.Now()}); result.Error != nil {
		return result.Error
	}

	return nil
}
