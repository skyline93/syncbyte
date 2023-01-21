package scheduling

import (
	"encoding/json"

	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"gorm.io/datatypes"
)

func ScheduleBackup(policyID uint) (jobID uint, err error) {
	var jobType = "backup"

	pl := backup.GetPolicy(policyID)

	tx := repository.Repo.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	bs := repository.BackupSet{
		IsValid:    false,
		ResourceID: pl.ResourceID,
		Retention:  pl.Retention,
	}

	if result := tx.Create(&bs); result.Error != nil {
		return 0, result.Error
	}

	v, err := json.Marshal(pl.Resource.Args)
	if err != nil {
		return 0, err
	}

	sj := repository.ScheduledJob{
		JobType:      jobType,
		Status:       "queued",
		ResourceID:   pl.ResourceID,
		ResourceType: pl.Resource.Type,
		Args:         datatypes.JSON(v),

		BackupSetID: bs.ID,
	}

	if result := tx.Create(&sj); result.Error != nil {
		return 0, result.Error
	}

	return sj.ID, nil
}
