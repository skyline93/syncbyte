package database

import "gorm.io/gorm"

type Database struct {
	*gorm.DB
}

func New(opts *Options) (*Database, error) {
	dia := Dialector(opts)

	db, err := gorm.Open(dia, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
