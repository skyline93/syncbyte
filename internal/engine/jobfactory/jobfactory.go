package jobfactory

import (
	"encoding/json"

	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"github.com/skyline93/syncbyte-go/internal/engine/syncbyte"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func ScheduleBackup(db *gorm.DB, resourceID uint) (jobID uint, err error) {
	var jobType = "backup"

	pl, err := syncbyte.GetPolicyByResource(db, resourceID)
	if err != nil {
		return 0, err
	}

	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	bs := syncbyte.BackupSet{
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
