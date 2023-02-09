package config

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/skyline93/syncbyte-go/pkg/database"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Core     CoreConfig       `mapstructure:"core"`
	Database database.Options `mapstructure:"database"`
}

type CoreConfig struct {
	ListenAddress    string `mapstructure:"listen_addr"`
	GrpcServerAddr   string `mapstructure:"grpc_server_addr"`
	LogPath          string `mapstructure:"log_path"`
	LogLevel         string `mapstructure:"log_level"`
	MongodbUri       string `mapstructure:"mongodb_uri"`
	MaxConcurrentNum int    `mapstructure:"max_concurrent_num"`
	JobSchedulerCron string `mapstructure:"job_scheduler_cron"`
}

func InitConfig() {
	fileName := "engine.toml"

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

	viper.SetDefault("core.grpc_server_addr", "127.0.0.1:50051")
	viper.SetDefault("core.log_path", "/var/syncbyte/log")
	viper.SetDefault("core.log_level", "debug")
	viper.SetDefault("core.mongodb_uri", "mongodb://mongoadmin:123456@127.0.0.1:27017/?maxPoolSize=20&w=majority")
	viper.SetDefault("core.max_concurrent_num", 64)
	viper.SetDefault("core.job_scheduler_cron", "* * * * *")
	viper.SetDefault("core.listen_addr", "0.0.0.0:50052")
	viper.SetDefault("database.type", "postgresql")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "syncbyte")
	viper.SetDefault("database.password", "123456")
	viper.SetDefault("database.dbname", "syncbyte")

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
