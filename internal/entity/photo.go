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

	Name       string `gorm:"size:200;index;"`
	FileName   string `gorm:"size:300;index;"`
	FileHash   string `gorm:"size:128;index" json:"file_hash"`
	FileSize   int64  `json:"file_size"`
	FileType   string `gorm:"size:16" json:"file_type"`
	IsImported bool   `gorm:"default:false"`

	UserID  uint
	AlbumID uint
	Albums  []Album `gorm:"many2many:photos_albums;" yaml:"-"`
}

func (Photo) TableName() string {
	return "photos"
}

func (s *Photo) Create(albumID uint) (*Photo, error) {
	photo := &Photo{
		Name:       s.Name,
		FileName:   s.FileName,
		FileHash:   s.FileHash,
		FileSize:   s.FileSize,
		FileType:   s.FileType,
		IsImported: false,
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

func FindUnimportedPhotosByAlbum(id uint) ([]Photo, error) {
	var album Album

	if err := Db().Preload("Photos", "is_imported = ?", false).First(&album, id).Error; err != nil {
		return nil, err
	}

	return album.Photos, nil
}

func (s *Photo) SetIsImported(isImported bool) error {
	if err := Db().Model(&Photo{}).Where("id = ?", s.ID).Update("is_imported", isImported).Error; err != nil {
		return err
	}

	return nil
}
