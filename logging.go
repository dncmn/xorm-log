package xorm_log

import (
	"github.com/go-xorm/core"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	logger     = logrus.New()
	xormLogger = &XormLogger{}
	sllog      = &DiscardLogger{}
)

// default log options
const (
	DEFAULT_LOG_PREFIX = "[xorm]"
	DEFAULT_LOG_FLAG   = log.Ldate | log.Lmicroseconds
	DEFAULT_LOG_LEVEL  = DebugLevel
)

func (*XormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Print(v[:])
	}
	if v[0] == "log" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

// GetGormLogger get gorm logger
func GetXormLogger() *XormLogger {
	return xormLogger
}

// DiscardLogger don't log implementation for core.ILogger
type DiscardLogger struct {
	level   core.LogLevel
	showSQL bool
}

type XormLogger struct{}

var _ core.ILogger = &DiscardLogger{}

func GetLogger() *DiscardLogger {
	return sllog
}

// GetLogrusLogger get logrus logger
func GetLogrusLogger() *logrus.Logger {
	return logger
}

// Debug logs a message with debug log level
func (d *DiscardLogger) Debug(v ...interface{}) {
	logger.Debug(v...)
}

// DDebugWithF logs a message with Debug log level
func (d *DiscardLogger) Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Error empty implementation
func (d *DiscardLogger) Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf empty implementation
func (d *DiscardLogger) Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Info empty implementation
func (d *DiscardLogger) Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof empty implementation
func (d *DiscardLogger) Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Warn empty implementation
func (d *DiscardLogger) Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Warnf empty implementation
func (d *DiscardLogger) Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
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
