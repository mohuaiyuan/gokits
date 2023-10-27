package conf

import (
	"errors"
	"github.com/spf13/viper"
)

var (
	Conf         *viper.Viper
	ErrNoSupport = errors.New("not support")
)

type Config struct {
	Local   string
	Url     string
	IsWatch bool
}

func Init(c *Config) error {
	if c.Local != "" {
		return local(c)
	}
	if c.Url != "" {
		return remote(c)
	}
	return ErrNoSupport
}

func local(c *Config) (err error) {
	Conf = viper.New()
	Conf.SetConfigFile(c.Local)
	if err = Conf.ReadInConfig(); err != nil {
		return
	}
	if c.IsWatch {
		Conf.WatchConfig()
	}
	return
}

func remote(c *Config) (err error) {
	return ErrNoSupport
}
