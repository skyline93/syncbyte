package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Options struct {
	Type     DBType `mapstructrue:"type"`
	Host     string `mapstructrue:"host"`
	Port     int    `mapstructrue:"port"`
	User     string `mapstructrue:"user"`
	Password string `mapstructrue:"password"`
	DbName   string `mapstructrue:"dbname"`
	Extra    string `mapstructrue:"extra"`
}

type DBType string

const (
	PostgreSQL DBType = "postgresql"
	MySQL      DBType = "mysql"
	SQLite     DBType = "sqlite"
)

func (opts *Options) Dsn() (dsn string) {
	switch opts.Type {
	case MySQL:
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			opts.User, opts.Password, opts.Host, opts.Port, opts.DbName,
		)

		if opts.Extra != "" {
			dsn = fmt.Sprintf("%s?%s", dsn, opts.Extra)
		}
	case PostgreSQL:
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d",
			opts.Host, opts.User, opts.Password, opts.DbName, opts.Port,
		)
		if opts.Extra != "" {
			dsn = fmt.Sprintf("%s %s", dsn, opts.Extra)
		}
	case SQLite:
		dsn = opts.DbName
	default:
		panic("unsupported database type")
	}

	return dsn
}

func Dialector(opts *Options) gorm.Dialector {
	switch opts.Type {
	case MySQL:
		return mysql.Open(opts.Dsn())
	case PostgreSQL:
		return postgres.Open(opts.Dsn())
	case SQLite:
		return sqlite.Open(opts.Dsn())
	default:
		return nil
	}
}
