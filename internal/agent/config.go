package agent

import (
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
	Core    CoreConfig    `mapstructure:"core"`
	Storage StorageConfig `mapstructure:"storage"`
}

type CoreConfig struct {
	GrpcAddr string `mapstructure:"grpc_addr"`
	LogPath  string `mapstructure:"log_path"`
	LogLevel string `mapstructure:"log_level"`
}

type StorageConfig struct {
	Type                StorageType `mapstructure:"type"`
	NASVolumeMountPoint string      `mapstructure:"nas_volume_mountpoint"`
}

func InitConfig() {
	fileName := "agent.toml"

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	fileExt := path.Ext(fileName)
	fileSuffix := strings.TrimSuffix(fileName, fileExt)

	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Join(home, ".config"))
	viper.AddConfigPath("/etc/syncbyte")

	viper.SetConfigName(fileSuffix)
	viper.SetConfigType(fileExt[1:])

	viper.SetDefault("core.grpc_addr", "0.0.0.0:50051")
	viper.SetDefault("core.log_path", "/var/syncbyte/log")
	viper.SetDefault("core.log_level", "debug")
	viper.SetDefault("storage.type", "nas")
	viper.SetDefault("storage.nas_volume_mountpoint", "/var/syncbyte/data")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	Conf = &Config{}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
}
