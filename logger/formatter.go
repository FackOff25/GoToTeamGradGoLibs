package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
)

type Formatter struct {
	NoColors   bool
	TimeFormat string
}

// Format an log entry
func (f Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)

	// output buffer
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "[%s] ", time.Now().Format(f.TimeFormat))

	// write level
	level := strings.ToUpper(entry.Level.String())

	if !f.NoColors {
		fmt.Fprintf(b, "\x1b[%dm", levelColor)
	}

	b.WriteString("[")
	b.WriteString(level[:4])
	b.WriteString("] ")

	if !f.NoColors {
		b.WriteString("\x1b[0m")
	}
	/*
		if prefix, ok := entry.Data["prefix"]; ok {
			fmt.Fprintf(b, "[%s] ", prefix)
		}

		if rid, ok := entry.Data["rid"]; ok {
			var userIds []string
			for _, idType := range []string{"login", "email", "uid"} {
				if val, ok := entry.Data[idType]; ok {
					userIds = append(userIds, val.(string))
				}
			}
			fmt.Fprintf(b, "[%s] +%.6f <%s> ", rid.(string), time.Since(entry.Time).Seconds(), strings.Join(userIds, ":"))
		}
	*/

	// write message
	b.WriteString(entry.Message)

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel:
		return colorGray
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}
