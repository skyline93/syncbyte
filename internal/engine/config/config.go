package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/skyline93/syncbyte-go/pkg/database"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	GrpcServerAddr string           `mapstructure:"grpc_server_addr"`
	MongodbUri     string           `mapstructure:"mongodb_uri"`
	Database       database.Options `mapstructure:"database"`
}

func InitConfig() {
	fileName := "engine.toml"

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(home, fileName)
	fileExt := path.Ext(fileName)
	fileSuffix := strings.TrimSuffix(fileName, fileExt)

	viper.AddConfigPath(home)
	viper.SetConfigName(fileSuffix)
	viper.SetConfigType(fileExt[1:])

	viper.SetDefault("grpc_server_addr", "127.0.0.1:50051")
	viper.SetDefault("mongodb_uri", "mongodb://mongoadmin:123456@127.0.0.1:27017/?maxPoolSize=20&w=majority")
	viper.SetDefault("database.type", "postgresql")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "syncbyte")
	viper.SetDefault("database.password", "123456")
	viper.SetDefault("database.dbname", "syncbyte")

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		if err := viper.SafeWriteConfig(); err != nil {
			panic(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Conf = &Config{}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
}
