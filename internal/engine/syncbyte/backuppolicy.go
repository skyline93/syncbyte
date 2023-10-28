package syncbyte

import (
	"gorm.io/gorm"
)

type BackupPolicy struct {
	gorm.Model
	ResourceID uint
	Retention  int
	Status     string
}

func GetPolicyByResource(db *gorm.DB, resourceID uint) (*BackupPolicy, error) {
	pl := BackupPolicy{}
	if result := db.Where("resource_id = ?", resourceID).First(&pl); result.Error != nil {
		return nil, result.Error
	}

	return &pl, nil
}
