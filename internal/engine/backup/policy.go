package backup

import (
	"encoding/json"

	"github.com/skyline93/syncbyte-go/internal/engine/repository"
	"gorm.io/datatypes"
)

type PolicyStatus string

const (
	Enabled  PolicyStatus = "enabled"
	Disabled PolicyStatus = "disabled"
)

type ResourceType string

const (
	NAS      ResourceType = "nas"
	Database ResourceType = "database"
)

type Resource struct {
	Name string
	Type string
	Args interface{}
}

type NasResourceArgs struct {
	Dir string `json:"dir"`
}

func CreatePolicy(resource Resource, retention int) (policyID uint, err error) {
	tx := repository.Repo.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	v, err := json.Marshal(resource.Args)
	if err != nil {
		return 0, err
	}

	res := repository.Resource{
		Name: resource.Name,
		Type: resource.Type,
		Args: datatypes.JSON(v),
	}

	if result := tx.Create(&res); result.Error != nil {
		return 0, result.Error
	}

	pl := repository.BackupPolicy{
		ResourceID: res.ID,
		Retention:  retention,
		Status:     string(Enabled),
	}

	if result := tx.Create(&pl); result.Error != nil {
		return 0, result.Error
	}

	return pl.ID, nil
}

type Policy struct {
	ID         uint
	Retention  int
	Status     string
	ResourceID uint
	Resource   Resource
}

func GetPolicy(policyID uint) *Policy {
	pl := repository.BackupPolicy{}
	if result := repository.Repo.First(&pl, policyID); result.Error != nil {
		return nil
	}

	res := repository.Resource{}
	if result := repository.Repo.First(&res, pl.ID); result.Error != nil {
		return nil
	}

	return &Policy{
		ID:         pl.ID,
		Retention:  pl.Retention,
		Status:     pl.Status,
		ResourceID: pl.ResourceID,
		Resource: Resource{
			Name: res.Name,
			Type: res.Type,
			Args: res.Args,
		},
	}
}
