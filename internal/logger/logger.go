package logger

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

type CallerHook struct{}

func (hook *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(1); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["source"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
	}

	return nil
}

type TextFormatter struct {
	logrus.TextFormatter
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%v] [%v] [%v] %v\n", entry.Level, entry.Time.Format("2006-01-02 15:04:05"), entry.Data["source"], entry.Message)), nil
}

func NewTextLogger(filename string, maxAge time.Duration) (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &TextFormatter{}
	log.AddHook(&CallerHook{})
	if filename == "" || filename == "stdout" {
		return log, nil
	}

	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	out, err := rotatelogs.New(
		abs+".%Y%m%d%H",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithLinkName(abs),
		rotatelogs.WithMaxAge(maxAge),
	)
	if err != nil {
		return nil, err
	}
	log.Out = out

	return log, nil
}

func NewJsonLogger(filename string, maxAge time.Duration) (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	if filename == "" || filename == "stdout" {
		return log, nil
	}

	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	out, err := rotatelogs.New(
		abs+".%Y%m%d%H",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithLinkName(abs),
		rotatelogs.WithMaxAge(maxAge),
	)
	if err != nil {
		return nil, err
	}
	log.Out = out

	return log, nil
}

func NewTextLoggerWithViper(v *viper.Viper) (*logrus.Logger, error) {
	return NewTextLogger(v.GetString("filename"), v.GetDuration("maxAge"))
}

func NewJsonLoggerWithViper(v *viper.Viper) (*logrus.Logger, error) {
	return NewJsonLogger(v.GetString("filename"), v.GetDuration("maxAge"))
}
