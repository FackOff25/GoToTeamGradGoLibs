package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Options struct {
	Name      string
	LogLevel  log.Level
	LogFormat string
	Out       io.Writer
}

func InitEx(opts Options) {
	log.SetLevel(opts.LogLevel)
	if opts.Out == nil {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(opts.Out)
	}

	var formatter log.Formatter = &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	if opts.LogFormat == "plain" {
		formatter = &Formatter{
			NoColors:   true,
			TimeFormat: "2006-01-02 15:04:05",
		}
	}
	log.SetFormatter(formatter)
}

func Init(name string) {
	InitEx(Options{
		Name:      name,
		LogLevel:  log.TraceLevel,
		LogFormat: "plain",
	})
}
