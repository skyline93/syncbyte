package config

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Core CoreConfig `mapstructure:"core"`
}

type CoreConfig struct {
	ServerAddress string `mapstructure:"server_addr"`
	LogLevel      string `mapstructure:"log_level"`
}

func InitConfig() {
	fileName := "syncbytectl.toml"

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	fileExt := path.Ext(fileName)
	fileSuffix := strings.TrimSuffix(fileName, fileExt)

	viper.AddConfigPath(filepath.Join(home, ".config"))

	viper.SetConfigName(fileSuffix)
	viper.SetConfigType(fileExt[1:])

	viper.SetDefault("core.server_addr", "127.0.0.1:50051")
	viper.SetDefault("core.log_level", "debug")

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
