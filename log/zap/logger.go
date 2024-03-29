package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

// 创建默认logger对象 方便外部直接调用包方法时使用
var logger *Logger = New(os.Stderr, DebugLevel, WithCaller(true), AddCallerSkip(1))

// 定义日志结构体
type Logger struct {
	log   *zap.Logger
	level *zap.AtomicLevel
}

// 定义Logger构造方法
func New(out io.Writer, level Level, opts ...Option) *Logger {
	if out == nil {
		out = os.Stderr
	}
	logLevel := zap.NewAtomicLevelAt(level)
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(out),
		logLevel)
	return &Logger{log: zap.New(core, opts...), level: &logLevel}
}

// 设置输出级别
func (l *Logger) SetLevel(level Level) {
	if l.level != nil {
		l.level.SetLevel(level)
	}
}

// Logger对象相关方法
func (l *Logger) Debug(msg string, fields ...Field) {
	l.log.Debug(msg, fields...)
}
func (l *Logger) Info(msg string, fields ...Field) {
	l.log.Info(msg, fields...)
}
func (l *Logger) Warn(msg string, fields ...Field) {
	l.log.Warn(msg, fields...)
}
func (l *Logger) Error(msg string, fields ...Field) {
	l.log.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.log.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.log.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.log.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.log.Sync()
}

// 返回包内置的logger对象
func Default() *Logger {
	return logger
}

// 替换包内置的logger对象
func ReplaceDefault(l *Logger) {
	logger = l
}

// 包方法 使用内置的logger对象完成操作
func SetLevel(level Level) {
	logger.SetLevel(level)
}

func Debug(msg string, fields ...Field) {
	logger.Debug(msg, fields...)
}
func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}
func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}
func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}
func DPanic(msg string, fields ...Field) {
	logger.DPanic(msg, fields...)
}
func Panic(msg string, fields ...Field) {
	logger.Panic(msg, fields...)
}
func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, fields...)
}

func Sync() error {
	return logger.Sync()
}
