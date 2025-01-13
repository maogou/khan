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
	nKey     string
	validate *validator.Validate
}

func NewGeweSdk(conf *config.Config, client *resty.Client, validate *validator.Validate) *Gewe {
	client.SetTimeout(conf.Sdk.TimeOut * time.Second)
	client.BaseURL = conf.Sdk.Pact

	return &Gewe{
		client:   client,
		config:   conf,
		token:    conf.Sdk.Token,
		appId:    conf.Sdk.AppId,
		uuid:     conf.Sdk.UuId,
		validate: validate,
	}
}

func (g *Gewe) SetToken(token string) *Gewe {
	g.token = token
	return g
}

func (g *Gewe) GetToken() string {
	return g.token
}

func (g *Gewe) SetAppId(appId string) *Gewe {
	g.appId = appId

	return g
}

func (g *Gewe) GetAppId() string {
	return g.appId
}

func (g *Gewe) SetUuId(uuid string) *Gewe {
	g.uuid = uuid

	return g
}
func (g *Gewe) GetUuId() string {
	return g.uuid
}

func (g *Gewe) SetNKey(key string) *Gewe {
	g.nKey = key

	return g
}
func (g *Gewe) GetNKey() string {
	return g.nKey
}

func (g *Gewe) Config() *config.Config {
	return g.config
}

func (g *Gewe) Client() *resty.Client {
	return g.client
}

func (g *Gewe) Validate() *validator.Validate {
	return g.validate
}
