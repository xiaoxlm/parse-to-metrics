package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Log struct {
	Name  string
	Level string
}

func (log *Log) SetDefaults() *Log {
	if log.Name == "" {
		log.Name = "parse-to-metrics"
	}

	if log.Level == "" {
		log.Level = "DEBUG"
	}

	return log
}

func (log *Log) Build() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier:  prettyfier,
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableHTMLEscape: true,
	})

	logrus.SetLevel(getLogLevel(log.Level))
	logrus.SetReportCaller(true)
	logrus.AddHook(NewServiceHook(log.Name))
	logrus.SetOutput(os.Stdout)
}

func prettyfier(f *runtime.Frame) (function string, file string) {
	return f.Function + " line:" + strconv.FormatInt(int64(f.Line), 10), ""
}

func getLogLevel(l string) logrus.Level {
	level, err := logrus.ParseLevel(strings.ToLower(l))
	if err == nil {
		return level
	}
	return logrus.InfoLevel
}
