package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

type Config struct {
	// 共用配置
	RotateType string           // 轮转方式
	Filename   string           // 完整文件名
	MaxAge     int              // 保留旧日志文件的最大天数
	Lef        LevelEnablerFunc // 日志级别启用函数

	// 按时间轮转配置
	RotationTime time.Duration // 日志文件轮转时间

	// 按大小轮转配置
	MaxSize    int  // 日志文件最大大小（MB）
	MaxBackups int  // 保留日志文件的最大数量
	Compress   bool // 是否对日志文件进行压缩归档
	LocalTime  bool // 是否使用本地时间，默认 UTC 时间
}

// NewWithRotate 创建日志对象按轮转方式 io.Writer
func NewWithRotate(config *Config, opts ...Option) *Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	lv := zap.LevelEnablerFunc(config.Lef)
	var writer io.Writer
	if config.RotateType == "byTime" {
		writer = NewRotateByTime(config)
	} else {
		writer = NewRotateBySize(config)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(os.Stdout), lv),
		zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(writer), lv),
	)

	return &Logger{log: zap.New(core, opts...)}

}

func NewRotateByTime(cfg *Config) io.Writer {
	opts := []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge) * time.Hour * 24),
		rotatelogs.WithRotationTime(cfg.RotationTime),
		rotatelogs.WithLinkName(cfg.Filename),
	}
	if !cfg.LocalTime {
		rotatelogs.WithClock(rotatelogs.UTC)
	}
	filename := strings.SplitN(cfg.Filename, ".", 2)
	w, _ := rotatelogs.New(filename[0]+"_%Y%m%d."+filename[1], opts...)
	return w
}

func NewRotateBySize(cfg *Config) io.Writer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	})
}
