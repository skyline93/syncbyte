package syncbyte

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name string `gorm:"unique"`
	Type string
	Args datatypes.JSON
}
