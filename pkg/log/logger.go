package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hlhgogo/athena/pkg/config"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

const (
	Type                   = "app"
	DefaultTimestampFormat = "2006-01-02 15:04:05"
)

var (
	log *logrus.Logger
)

// Trace trace log
func Trace(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Trace(args...)
}

// Tracef info log
func Tracef(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Trace(args...)
}

// Debug info log
func Debug(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Debug(args...)
}

// Debugf info log
func Debugf(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Debug(args...)
}

// Info info log
func Info(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Info(args...)
}

// Infof 格式化输出info log
func Infof(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Info(fmt.Sprintf(format, args...))
}

// Warn warn log
func Warn(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Warn(args...)
}

// Warnf 格式化输出warn log
func Warnf(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Warn(fmt.Sprintf(format, args...))
}

// Error error log
func Error(args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Error(args...)
}

// Errorf error 格式化输出warn log
func Errorf(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"type": Type,
	}).Error(fmt.Sprintf(format, args...))
}

func GetGinLogIoWriter() io.Writer {
	writer, err := rotatelogs.New(
		config.Get().Logger.SavePath+"/api-%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Duration(config.Get().Logger.SaveDay)*24*time.Hour),       // Maximum file save time
		rotatelogs.WithRotationTime(time.Duration(config.Get().Logger.SaveDay)*24*time.Hour), // Log the cut interval
	)
	if err != nil {
		panic(err)
	}

	return writer
}

func GetProjectIoWriter() io.Writer {
	writer, err := rotatelogs.New(
		config.Get().Logger.SavePath+"/gin-%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Duration(config.Get().Logger.SaveDay)*24*time.Hour),       // Maximum file save time
		rotatelogs.WithRotationTime(time.Duration(config.Get().Logger.SaveDay)*24*time.Hour), // Log the cut interval
	)
	if err != nil {
		panic(err)
	}

	return writer
}

type LineFormatter struct {
	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string
}

// Format implement the Formatter interface
func (f *LineFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}

	var field = ""
	if entry.Data != nil {
		if b, err := json.Marshal(entry.Data); err == nil {
			field = string(b)
		}

	}

	b.WriteString(fmt.Sprintf("%s [%s] [%s] %s - %s %s\n", config.Get().App.Name, strings.ToUpper(entry.Level.String()), entry.Time.Format(timestampFormat),
		fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line), entry.Message, field))

	return b.Bytes(), nil
}

func Setup() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	log = logrus.New()

	log.SetReportCaller(true)
	log.SetFormatter(&LineFormatter{TimestampFormat: DefaultTimestampFormat})

	writer := GetProjectIoWriter()
	writers := []io.Writer{writer}

	// Local development output to the console
	if env == "local" {
		writers = append(writers, os.Stdout)
	}

	log.SetOutput(io.MultiWriter(writers...))
	log.SetFormatter(&LineFormatter{TimestampFormat: DefaultTimestampFormat})
	//log.SetFormatter(&log.JSONFormatter{TimestampFormat: DefaultTimestampFormat})

	// Set log level
	var level logrus.Level = logrus.TraceLevel
	switch config.Get().Logger.Level {
	case "trace":
		level = logrus.TraceLevel
	case "debug":
		level = logrus.DebugLevel
	case "warn":
		level = logrus.WarnLevel
	case "info":
		level = logrus.InfoLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	}
	log.SetLevel(level)
}
