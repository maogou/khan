package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	DefaultConfigPath = "config/robot.yaml"
)

type Config struct {
	Mode    string `yaml:"mode" validate:"required,oneof=debug test release"`
	Port    int    `yaml:"port" validate:"required"`
	TimeOut int    `yaml:"timeout" validate:"required,min=2"`
}

func MustLoadConfig(cfgFile string) Config {
	if len(cfgFile) == 0 {
		cfgFile = DefaultConfigPath
	}

	conf := viper.New()

	conf.SetConfigFile(cfgFile)
	conf.SetConfigType("yaml")

	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := conf.Unmarshal(&config); err != nil {
		panic(err)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		panic(err)
	}

	return config
}
