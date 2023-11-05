package syncbyte

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = gormDB.AutoMigrate(
		&Resource{},
		&BackupPolicy{},
		&ScheduledJob{},
		&BackupSet{},
	)
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
