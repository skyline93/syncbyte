package repository

import (
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/pkg/database"
)

var Repo *Repository

func InitRepository() {
	var err error

	Repo, err = New(&config.Conf.Database)
	if err != nil {
		panic(err)
	}

	if err = Repo.AutoMigrate(
		&DBResource{},
		&Resource{},
		&S3Backend{},
		&BackupJob{},
		&BackupSet{},
		&RestoreJob{},
		&RestoreDBResource{},
		&BackupPolicy{},
		&Agent{},
		&ScheduledJob{},
	); err != nil {
		panic(err)
	}
}

type Repository struct {
	*database.Database
}

func New(opts *database.Options) (*Repository, error) {
	db, err := database.New(opts)
	if err != nil {
		return nil, err
	}

	return &Repository{Database: db}, nil
}
