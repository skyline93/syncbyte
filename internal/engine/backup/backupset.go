package backup

import (
	"time"

	"github.com/skyline93/syncbyte-go/internal/engine/repository"
)

type BackupSet struct {
	ID         uint
	IsValid    bool
	Size       int64
	BackupTime time.Time
	ResourceID uint
	Retention  int
}

func GetBackupSet(bsetID uint) (bset *BackupSet, err error) {
	item := repository.BackupSet{}

	if result := repository.Repo.First(&item, bsetID); result.Error != nil {
		return nil, result.Error
	}

	return &BackupSet{
		ID:         item.ID,
		IsValid:    item.IsValid,
		Size:       item.Size,
		BackupTime: item.BackupTime,
		ResourceID: item.ResourceID,
		Retention:  item.Retention,
	}, nil
}

func (s *BackupSet) SetBackupTime(bt time.Time) error {
	if result := repository.Repo.Model(&repository.BackupSet{}).Where("id = ?", s.ID).Update("backup_time", bt); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *BackupSet) SetAvailability(isValid bool) error {
	if result := repository.Repo.Model(&repository.BackupSet{}).Where("id = ?", s.ID).Update("is_valid", isValid); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *BackupSet) SetSize(size int64, scanSize int64) error {
	if result := repository.Repo.Model(&repository.BackupSet{}).Where("id = ?", s.ID).Updates(repository.BackupSet{Size: size, ScanSize: scanSize}); result.Error != nil {
		return result.Error
	}
	return nil
}
