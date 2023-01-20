package logging

import (
	"io"

	log "github.com/sirupsen/logrus"
)

func LogrusLevel(level string) log.Level {
	switch level {
	case "fatal":
		return log.FatalLevel
	case "error":
		return log.ErrorLevel
	case "warn":
		return log.WarnLevel
	case "info":
		return log.InfoLevel
	case "debug":
		return log.DebugLevel
	default:
		panic("logger level unknow")
	}
}

type FormatterHook struct {
	Writer    io.Writer
	LogLevels []log.Level
	Formatter log.Formatter
}

func (hook *FormatterHook) Fire(entry *log.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *FormatterHook) Levels() []log.Level {
	return hook.LogLevels
}
