package xorm_log

import (
	"github.com/go-xorm/core"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	logger = logrus.New()
	sllog  = &DiscardLogger{}
)

type LogLevel int

const (
	// !nashtsai! following level also match syslog.Priority value
	LOG_DEBUG LogLevel = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERR
	LOG_OFF
	LOG_UNKNOWN
)

// default log options
const (
	DEFAULT_LOG_PREFIX = "[xorm]"
	DEFAULT_LOG_FLAG   = log.Ldate | log.Lmicroseconds
	DEFAULT_LOG_LEVEL  = DebugLevel
)

// GetGormLogger get gorm logger

// DiscardLogger don't log implementation for core.ILogger
type DiscardLogger struct {
	entry   *logrus.Entry
	level   core.LogLevel
	showSQL bool
}

//var _ core.ILogger = &DiscardLogger{}

func GetLogger() *DiscardLogger {
	return sllog
}

// GetLogrusLogger get logrus logger
func GetLogrusLogger() *logrus.Logger {
	return logger
}

// Error implement core.ILogger
func (s *DiscardLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.entry.Error(v...)
	}
	return
}

// Errorf implement core.ILogger
func (s *DiscardLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.entry.Errorf(format, v...)
	}
	return
}

// Debug implement core.ILogger
func (s *DiscardLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.entry.Debug(v...)
	}
	return
}

// Debugf implement core.ILogger
func (s *DiscardLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.entry.Debugf(format, v...)
	}
	return
}

// Info implement core.ILogger
func (s *DiscardLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.entry.Info(v...)
	}
	return
}

// Infof implement core.ILogger
func (s *DiscardLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.entry.Infof(format, v...)
	}
	return
}

// Warn implement core.ILogger
func (s *DiscardLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.entry.Warn(v...)
	}
	return
}

// Warnf implement core.ILogger
func (s *DiscardLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.entry.Warnf(format, v...)
	}
	return
}

func Init(config FileLogConfig) {
	os.MkdirAll(config.Path, 0777)
	hook, err := newFileHook(config)
	if err != nil {
		logger.Error(err)
	}
	logger.Level = logrus.DebugLevel
	logger.AddHook(hook)
}

// Level implement core.ILogger
func (s *DiscardLogger) Level() core.LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *DiscardLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

// ShowSQL implement core.ILogger
func (s *DiscardLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *DiscardLogger) IsShowSQL() bool {
	return s.showSQL
}
