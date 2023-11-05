package database

import (
	"github.com/mitchellh/mapstructure"
)

type MySQL struct {
	DbType string `mapstructure:"dbType"`
	DbName string `mapstructure:"dbName"`
	Port   int    `mapstructure:"port"`
}

func NewMySQL(dbName string, port int) *MySQL {
	return &MySQL{DbName: dbName, Port: port, DbType: DbTypeMySQL}
}

func (m *MySQL) GetAttr() (map[string]interface{}, error) {
	var v map[string]interface{}

	err := mapstructure.Decode(m, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
