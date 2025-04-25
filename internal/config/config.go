package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	DefaultConfigPath = "config/khan.yaml"
)

type Config struct {
	Mode    string     `yaml:"mode" validate:"required,oneof=debug test release"`
	Port    int        `yaml:"port" validate:"required"`
	TimeOut int        `yaml:"timeout" validate:"required,min=2"`
	Sdk     *Sdk       `yaml:"sdk" validate:"required"`
	Redis   *RdbConfig `yaml:"redis" validate:"required"`
}

type Sdk struct {
	Collect   string        `yaml:"collect"`
	Callback  string        `yaml:"callback"`
	TimeOut   time.Duration `yaml:"timeout" validate:"required,min=5"`
	Gog7a6v8g string        `yaml:"gog7a6v8g" validate:"required"`
	Gog7a6v90 string        `yaml:"gog7a6v90" validate:"required"`
	Token     string        `yaml:"token"`
	AppId     string        `yaml:"appid"`
	UuId      string        `yaml:"uuid"`
	License   string        `yaml:"license" validate:"required"`
}

type RdbConfig struct {
	Addr     string `yaml:"addr" validate:"required"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db" validate:"max=15"`
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
