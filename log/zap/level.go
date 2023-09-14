package log

import "go.uber.org/zap/zapcore"

type Level = zapcore.Level

// 声明日志级别常量
const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)
