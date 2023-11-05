package job

import (
	"log"

	"github.com/skyline93/syncbyte/syncbyte"
	"gorm.io/gorm"
)

type BackupJob struct {
	syncbyte.ScheduledJob
}

func (j *BackupJob) Run(db *gorm.DB, resourceID uint, retention int) (err error) {
	defer func() {
		if err != nil {
			j.Fail(db, err.Error())
			return
		}

		j.Success(db)
	}()

	if err = j.Start(db); err != nil {
		return err
	}

	var bset *syncbyte.BackupSet
	log.Printf("create backup set")
	bset, err = syncbyte.CreateBackupSet(db, resourceID, retention)
	if err != nil {
		return err
	}

	log.Printf("create backup set completed, id: %d", bset.ID)

	return
}
