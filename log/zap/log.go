package log

import (
	"github.com/mohuaiyuan/gokits/pkg/conf"
	"time"
)

var (
	_maxAge     = 10
	_maxSize    = 2
	_maxBackups = 30
)

func loadLogConfig() (*Config, error) {
	config := new(Config)
	rotateType := conf.Conf.GetString("log.rotateType")
	filename := conf.Conf.GetString("log.filename")
	maxAge := conf.Conf.GetInt("log.maxAge")
	level := conf.Conf.GetString("log.level")
	rotationTime := conf.Conf.GetInt("log.rotationTime")
	maxSize := conf.Conf.GetInt("log.maxSize")
	maxBackups := conf.Conf.GetInt("log.maxBackups")
	compress := conf.Conf.GetBool("log.compress")
	localTime := conf.Conf.GetBool("log.localTime")

	if rotateType == "bySize" {
		config.RotateType = "bySize"
	} else {
		config.RotateType = "byTime"
	}

	if filename == "" {
		config.Filename = "log.txt"
	} else {
		config.Filename = filename
	}

	if maxAge == 0 {
		config.MaxAge = _maxAge
	} else {
		config.MaxAge = maxAge
	}

	if level == "" {
		level = "debug"
	}
	switch level {
	case "debug":
		config.Lef = SetLevelFunc(DebugLevel)
	case "info":
		config.Lef = SetLevelFunc(InfoLevel)
	case "warn":
		config.Lef = SetLevelFunc(WarnLevel)
	case "error":
		config.Lef = SetLevelFunc(ErrorLevel)
	case "dPanic":
		config.Lef = SetLevelFunc(DPanicLevel)
	case "panic":
		config.Lef = SetLevelFunc(PanicLevel)
	case "fatal":
		config.Lef = SetLevelFunc(FatalLevel)
	}
	if rotationTime == 0 {
		config.RotationTime = time.Hour * 24 * 1
	} else {
		config.RotationTime = time.Hour * 24 * time.Duration(rotationTime)
	}
	if maxSize == 0 {
		config.MaxSize = _maxSize
	} else {
		config.MaxSize = maxSize
	}
	if maxBackups == 0 {
		config.MaxBackups = _maxBackups
	} else {
		config.MaxBackups = maxBackups
	}
	config.Compress = compress
	config.LocalTime = localTime
	return config, nil
}

func init() {
	config, _ := loadLogConfig()
	logger = NewWithRotate(config, WithCaller(true), AddCallerSkip(1))
	ReplaceDefault(logger)
}
