package config

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/fsnotify/fsnotify"

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
	Sdk     *Sdk   `yaml:"sdk" validate:"required"`
}

type Sdk struct {
	Collect  string        `yaml:"collect" validate:"required"`
	Callback string        `yaml:"callback"`
	TimeOut  time.Duration `yaml:"timeout" validate:"required,min=5"`
	Pact     string        `yaml:"pact" validate:"required"`
	Long     string        `yaml:"long" validate:"required"`
	Token    string        `yaml:"token"`
	AppId    string        `yaml:"appid"`
	UuId     string        `yaml:"uuid"`
	License  string        `yaml:"license" validate:"required"`
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

	go func(conf *viper.Viper) {
		watchConfig(conf, cfgFile)
	}(conf)

	return config
}

func watchConfig(conf *viper.Viper, cfgFile string) {
	conf.WatchConfig()

	conf.OnConfigChange(
		func(e fsnotify.Event) {
			log.Info().Str("file", cfgFile).Msg("配置文件已修改")
			if err := conf.ReadInConfig(); err != nil {
				log.Error().Err(err).Str("file", cfgFile).Msg("重新读取配置文件失败")
			} else {
				log.Info().Str("file", cfgFile).Msg("配置文件重新加载成功")
				printAllConfig(conf)
			}
		},
	)
}

func printAllConfig(conf *viper.Viper) {
	allSettings := conf.AllSettings()
	for key, value := range allSettings {
		log.Info().Str("key", key).Interface("value", value).Msg("当前配置项:")
	}
}
