package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var loggerMap = LoggerMap{}

type LoggerMap struct {
	sync.Map
	rootPath string
}

func SetupLogger(rootPath string) {
	loggerMap.rootPath = rootPath
}

func GetLogger(name string) *log.Logger {
	v, ok := loggerMap.Load(name)
	if ok {
		return v.(*log.Logger)
	}

	if loggerMap.rootPath == "" {
		loggerMap.rootPath = "logs"
	}

	hook := lumberjack.Logger{
		Filename:  filepath.Join(loggerMap.rootPath, fmt.Sprintf("%s.log", name)),
		MaxSize:   1024,
		MaxAge:    365,
		Compress:  true,
		LocalTime: true,
	}

	loggerEntity := log.New()
	loggerEntity.SetFormatter(&log.JSONFormatter{})
	loggerEntity.SetOutput(io.MultiWriter(&hook, os.Stdout))
	loggerEntity.SetLevel(log.DebugLevel)

	loggerMap.Store(name, loggerEntity)

	return loggerEntity
}
