package log

import (
	"context"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gonb/pkg/config"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Fields map[string]interface{}

// Define logrus alias
var (
	Logger          = NewLog()
	Log             = Logger
	Tracef          = Logger.Tracef
	Debugf          = Logger.Debugf
	Infof           = Logger.Infof
	Warnf           = Logger.Warnf
	Errorf          = Logger.Errorf
	Fatalf          = Logger.Fatalf
	Panicf          = Logger.Panicf
	Printf          = Logger.Printf
	Println         = Logger.Println
	Write           = Logger.Error
	SetOutput       = Logger.SetOutput
	WithFields      = Logger.WithFields
	WithField       = Logger.WithField
	SetReportCaller = Logger.SetReportCaller
	StandardLogger  = logrus.StandardLogger
	ParseLevel      = logrus.ParseLevel
	debug           int
)

// Define logger level
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// Log Logrus
type LoggerV2 struct {
	*logrus.Logger
}

// Entry logrus.Entry alias
type Entry = logrus.Entry

//
// Hook logrus.Hook alias
type Hook = logrus.Hook

type Level = logrus.Level

func NewLog() *LoggerV2 {
	loggerV2 := &LoggerV2{
		logrus.New(),
	}
	// 默认是json格式和终端输出
	format := config.GetStringOrDefault("log.format", "text")
	switch format {
	case "text":
		loggerV2.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05.000", FullTimestamp: true}
	case "json":
		loggerV2.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"}
	}
	// Logger.Out       = os.Stdout
	logPath := config.GetString("log.path")
	debug = config.GetIntOrDefault("log.debug", 0)
	loggerV2.SetFileOutWriter(logPath, "app", 7*24*time.Hour, 1*time.Hour)
	return loggerV2
}

// SetOutWriter  可以设置文件记录日志
func (l *LoggerV2) SetOutWriter(writer io.Writer) {
	l.Out = writer
}

func (l *LoggerV2) getFormat() logrus.Formatter {
	return l.Logger.Formatter
}

// IsExist util存在对应方法，避免循环引用
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}

// SetFileOutWriter 设置带文件并且切割的日志模式
func (l *LoggerV2) SetFileOutWriter(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	// pro模式不开控制台日志
	if debug == 0 {
		l.SetOutput(ioutil.Discard)
	}
	if IsExist(logPath) == false {
		_ = os.Mkdir(logPath, 0755)
	}
	path := fmt.Sprintf("%s/%s.log", logPath, logFileName)
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	l.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		},
		&logrus.JSONFormatter{},
	))
}

func (l *LoggerV2) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *LoggerV2) Tracef(format string, args ...interface{}) {
	l.Logger.Tracef(format, args...)
}

func (l *LoggerV2) Debugf(format string, args ...interface{}) {
	l.Logger.Tracef(format, args...)
}

func (l *LoggerV2) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *LoggerV2) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *LoggerV2) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}
func (l *LoggerV2) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}
func (l *LoggerV2) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}
func (l *LoggerV2) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

func (l *LoggerV2) WithFields(fields Fields) *Entry {
	return l.Logger.WithFields(logrus.Fields(fields))
}
func (l *LoggerV2) WithField(key string, value interface{}) *Entry {
	return l.Logger.WithField(key, value)
}

func (l *LoggerV2) SetOutput(output io.Writer) {
	l.Logger.SetOutput(output)
}

func (l *LoggerV2) SetReportCaller(reportCaller bool) {
	l.Logger.SetReportCaller(reportCaller)
}

// Write 为了兼容以前老的代码
func (l *LoggerV2) Write(flag string, err error, infos ...interface{}) {
	var data = make(Fields)
	data["flag"] = flag
	data["err"] = err
	data["info"] = infos
	if Logger.Level == logrus.FatalLevel {
		l.WithFields(data).Error(infos...)
	} else {
		l.WithFields(data).Info(infos...)
	}
}

func (l *LoggerV2) WriteMap(val map[string]interface{}, infos ...interface{}) {
	if Logger.Level == logrus.FatalLevel {
		l.WithFields(val).Error(infos...)
	} else {
		l.WithFields(val).Info(infos...)
	}
}

// SetLevel Set logger level
func (l *LoggerV2) SetLevel(level Level) {
	l.Level = level
}

// SetFormatter Set logger output format (json/text)
func (l *LoggerV2) SetFormatter(format string) {
	switch format {
	case "json":
		l.Formatter = new(logrus.JSONFormatter)
	case "text":
		l.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	default:
		l.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	}
}

// AddLogHook AddHook Add logger hook
func (l *LoggerV2) AddLogHook(hook Hook) {
	l.Logger.AddHook(hook)
}

// Define key
const (
	TraceIDKey  = "trace_id"
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
	TagKey      = "tag"
	StackKey    = "stack"
)

type (
	traceIDKey  struct{}
	userIDKey   struct{}
	userNameKey struct{}
	tagKey      struct{}
	stackKey    struct{}
)

func (*LoggerV2) NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func (*LoggerV2) FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (*LoggerV2) NewUserIDContext(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func (*LoggerV2) FromUserIDContext(ctx context.Context) uint64 {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}

func (*LoggerV2) NewUserNameContext(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func (*LoggerV2) FromUserNameContext(ctx context.Context) string {
	v := ctx.Value(userNameKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (*LoggerV2) NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

func (*LoggerV2) FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}

	return ""
}

func (*LoggerV2) NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

func (l *LoggerV2) FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}

	return nil
}

// WithContext Use context create entry
func (l *LoggerV2) WithContext(ctx context.Context) *Entry {
	fields := logrus.Fields{}

	if v := l.FromTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}

	if v := l.FromUserIDContext(ctx); v != 0 {
		fields[UserIDKey] = v
	}

	if v := l.FromUserNameContext(ctx); v != "" {
		fields[UserNameKey] = v
	}

	if v := l.FromTagContext(ctx); v != "" {
		fields[TagKey] = v
	}

	if v := l.FromStackContext(ctx); v != nil {
		fields[StackKey] = fmt.Sprintf("%+v", v)
	}

	return l.Logger.WithContext(ctx).WithFields(fields)
}
