package log

import "go.uber.org/zap/zapcore"

type Level = zapcore.Level

// 声明日志级别常量
const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	DPanicLevel = zapcore.DPanicLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)

type LevelEnablerFunc func(lvl Level) bool

func SetLevelFunc(level Level) LevelEnablerFunc {
	return func(lvl Level) bool {
		return lvl >= level
	}
}

func ErrorLevelFunc() LevelEnablerFunc {
	return func(lvl Level) bool {
		return lvl == ErrorLevel
	}
}
