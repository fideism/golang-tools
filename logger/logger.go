package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// Log 实例
var Log *logrus.Logger

func init() {
	Log = logrus.New()

	log.SetOutput(os.Stdout)

	if lv := os.Getenv("LOG_LEVEL"); "" != lv {
		if v, err := logrus.ParseLevel(lv); nil == err {
			Log.SetLevel(v)
		}
	}
}

// SetOption 设置参数
func SetOption(option ...LogstashFormatterOption)  {
	format := LogstashFormatter{}

	for _, o := range option {
		o(&format)
	}

	if len(format.Channel) == 0 {
		format.Channel = ""
	}

	Log.SetFormatter(&format)
}

// Logger 实例
type Logger struct {
	*logrus.Entry
}

// Entry 返回默认日志
func Entry() *Logger {
	return &Logger{
		Entry:logrus.NewEntry(Log),
	}
}

// WithError 重载error
func (entry *Logger) WithError(err error) *logrus.Entry {
	return entry.WithField("error", err.Error())
}

// NewEntry 返回默认日志
func NewEntry(option ...LogstashFormatterOption) *Logger {
	format := LogstashFormatter{}

	for _, o := range option {
		o(&format)
	}

	if len(format.Channel) == 0 {
		format.Channel = ""
	}

	Log.SetFormatter(&format)

	return &Logger{
		Entry:logrus.NewEntry(Log),
	}
}
