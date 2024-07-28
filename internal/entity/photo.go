package entity

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name     string  `gorm:"size:200;index;"`
	FileName string  `gorm:"size:300;index;"`
	FileHash string  `gorm:"size:128;index" json:"file_hash"`
	FileSize int64   `json:"file_size"`
	FileType string  `gorm:"size:16" json:"file_type"`
	Albums   []Album `gorm:"many2many:photos_albums;" yaml:"-"`
}

func (Photo) TableName() string {
	return "photos"
}

func (s *Photo) Create(albumID uint) (*Photo, error) {
	photo := &Photo{
		Name:     s.Name,
		FileName: s.FileName,
		FileHash: s.FileHash,
		FileSize: s.FileSize,
		FileType: s.FileType,
	}

	if err := Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(photo).Error; err != nil {
			return err
		}

		var album Album
		if err := tx.First(&album, albumID).Error; err != nil {
			logger.Debugf("err: %s", err)
			return err
		}

		if err := tx.Model(&album).Association("Photos").Append(photo); err != nil {
			logger.Debugf("err: %s", err)
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return photo, nil
}
