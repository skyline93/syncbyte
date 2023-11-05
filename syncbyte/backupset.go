package syncbyte

import (
	"time"

	"gorm.io/gorm"
)

type BackupSet struct {
	gorm.Model
	BackupTime *time.Time
	ResourceID uint
	Retention  int
}

func CreateBackupSet(db *gorm.DB, resourceID uint, retention int) (*BackupSet, error) {
	bs := BackupSet{ResourceID: resourceID, Retention: retention}
	if result := db.Create(&bs); result.Error != nil {
		return nil, result.Error
	}

	return &bs, nil
}
