package syncbyte

import (
	"errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BackupPolicy struct {
	gorm.Model
	ResourceID uint
	Resource   Resource
	Retention  int
}

func CreateBackupPolicy(db *gorm.DB, indentifier string, resourceType ResourceType, resourceAttr datatypes.JSONMap, retention int) (*BackupPolicy, error) {
	res, err := getResource(db, indentifier)
	if err != nil {
		return nil, err
	}

	if res == nil {
		res, err = createResource(db, indentifier, resourceType, resourceAttr)
		if err != nil {
			return nil, err
		}
	}

	bp := BackupPolicy{Resource: *res, Retention: retention}
	if result := db.Create(&bp); result.Error != nil {
		return nil, result.Error
	}

	return &bp, nil
}

func GetBackupPolicy(db *gorm.DB, policyID int) (*BackupPolicy, error) {
	var pl BackupPolicy
	if result := db.Where("id = ?", policyID).Preload("Resource").First(&pl); result.Error != nil {
		return nil, result.Error
	}

	return &pl, nil
}

func getResource(db *gorm.DB, indent string) (*Resource, error) {
	var res Resource
	if result := db.Where("identifier = ?", indent).First(&res); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &res, nil
}

func createResource(db *gorm.DB, indentifier string, resourceType ResourceType, attributes datatypes.JSONMap) (*Resource, error) {
	res := Resource{Identifier: indentifier, Type: resourceType, Attributes: attributes}
	if result := db.Create(&res); result.Error != nil {
		return nil, result.Error
	}

	return &res, nil
}
