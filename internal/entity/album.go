package entity

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name   string  `gorm:"size:200;index;"`
	Users  []User  `gorm:"many2many:users_albums;" yaml:"-"`
	Photos []Photo `gorm:"many2many:photos_albums;" yaml:"-"`
}

func (Album) TableName() string {
	return "albums"
}

func ExistsAlbum(name string) bool {
	result := Album{}
	if err := Db().Where("name = ?", name).First(&result).Error; err != nil {
		return false
	}

	return true
}

func (s *Album) Create(name string, userID uint) (*Album, error) {
	album := &Album{Name: s.Name}

	if err := Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(album).Error; err != nil {
			return err
		}

		var user User
		if err := tx.First(&user, userID).Error; err != nil {
			logger.Debugf("err: %s", err)
			return err
		}

		if err := tx.Model(&user).Association("Albums").Append(album); err != nil {
			logger.Debugf("err: %s", err)
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return album, nil
}

func FindAlbumsAll() ([]*Album, error) {
	var albums []*Album

	if err := Db().Find(&albums).Error; err != nil {
		return nil, err
	}

	return albums, nil
}

func FindAlbumByID(id uint) (*Album, error) {
	var album Album

	if err := Db().First(&album, id).Error; err != nil {
		return nil, err
	}

	return &album, nil
}

func (s *Album) Update(name string) error {
	if err := Db().Model(&Album{}).Update("name", name).Error; err != nil {
		return err
	}

	return nil
}

func (s *Album) Delete(id uint) error {
	if err := Db().Delete(&Album{}, id).Error; err != nil {
		return err
	}

	return nil
}
