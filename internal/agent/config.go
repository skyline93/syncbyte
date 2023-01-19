package agent

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var Conf *Config

type StorageType string

const (
	NAS StorageType = "nas"
	S3  StorageType = "s3"
)

type Config struct {
	GrpcAddr string  `mapstructure:"grpc_addr"`
	Storage  Storage `mapstructure:"storage"`
}

type Storage struct {
	Type                StorageType `mapstructure:"type"`
	NASVolumeMountPoint string      `mapstructure:"nas_volume_mountpoint"`
}

func InitConfig() {
	fileName := "agent.toml"

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

	viper.SetDefault("grpc_addr", "127.0.0.1:50051")
	viper.SetDefault("storage.type", "nas")
	viper.SetDefault("storage.nas_volume_mountpoint", "/var/syncbyte/data")

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
