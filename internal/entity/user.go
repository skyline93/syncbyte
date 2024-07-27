package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name     string `gorm:"size:200;index;"`
	Password string `gorm:"size:500"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) InvalidPassword(s string) bool {
	return u.Password == s
}

func FindUser(name string) *User {
	result := User{}
	if err := Db().Where("name = ?", name).First(&result).Error; err != nil {
		return nil
	}

	return &result
}
