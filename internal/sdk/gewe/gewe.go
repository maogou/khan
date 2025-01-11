package gewe

import (
	"smallBot/internal/config"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
)

type Gewe struct {
	client   *resty.Client
	config   *config.Config
	token    string
	appId    string
	uuid     string
	validate *validator.Validate
}

func NewGeweSdk(conf *config.Config, client *resty.Client, validate *validator.Validate) *Gewe {
	client.SetTimeout(conf.Sdk.TimeOut * time.Second)
	client.BaseURL = conf.Sdk.Url

	return &Gewe{
		client:   client,
		config:   conf,
		token:    conf.Sdk.Token,
		appId:    conf.Sdk.AppId,
		uuid:     conf.Sdk.UuId,
		validate: validate,
	}
}
