package db

import (
	"gonb/pkg/log"
	"gorm.io/gorm/logger"
	"time"
)

var (
	gormLogger = NewGormLogger()
	//newLogger  = GLog.New(
	//	log.Log,
	//	GLog.Config{
	//		SlowThreshold: 10 * time.Second, // 慢 SQL 阈值
	//		LogLevel:      GLog.Error,       // Log level
	//		Colorful:      false,            // 禁用彩色打印
	//	},
	//)
)

// GormWriter 定义自己的Writer
type GormWriter struct {
	gormLogger *log.LoggerV2
}

func (g *GormWriter) Printf(format string, v ...interface{})  {
	g.gormLogger.Infof(format, v...)
}

func NewGormLogger() logger.Interface {
	return logger.New(NewGormWriter(),
		logger.Config{
			//慢SQL阈值
			SlowThreshold: 2000 * time.Millisecond,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel: logger.Warn,
		})
}

func NewGormWriter() *GormWriter {
	return &GormWriter{gormLogger: log.Logger}
}
