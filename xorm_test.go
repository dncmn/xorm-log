package xorm_log

import "testing"

var (
	logging     *DiscardLogger
	xormLogging *XormLogger
)

func TestOrmLog(t *testing.T) {
	Init(FileLogConfig{
		Path:     "./test",
		Filename: "./test/app.log",
		MaxLines: 1000000,
		Maxsize:  1 << 28, //256 MB
		Daily:    true,
		MaxDays:  3,
		Rotate:   true,
		Level:    DebugLevel,
	})

	logging = GetLogger()
	xormLogging = GetXormLogger()
	logging.Info("logger success init.")
}

//func GetLogger() *sllog.Sllogger {
//	return logging
//}
//
//func GetGormLogger() *sllog.GormLogger {
//	return gormLogging
//}
