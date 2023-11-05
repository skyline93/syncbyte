package syncbyte

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ResourceType string

const (
	ResourceTypeDatabase   ResourceType = "database"
	ResourceTypeFileSystem ResourceType = "filesystem"
)

type Resource struct {
	gorm.Model
	Identifier string
	Type       ResourceType
	Attributes datatypes.JSONMap
}
