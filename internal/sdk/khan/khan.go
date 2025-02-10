package khan

import (
	"embed"
	"smallBot/internal/config"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
)

type Khan struct {
	client   *resty.Client
	config   *config.Config
	token    string
	appId    string
	uuid     string
	nKey     string
	validate *validator.Validate
	tpl      embed.FS
}

func NewKhanSdk(conf *config.Config, client *resty.Client, validate *validator.Validate, tpl embed.FS) *Khan {
	client.SetTimeout(conf.Sdk.TimeOut * time.Second)
	client.BaseURL = conf.Sdk.Pact

	return &Khan{
		client:   client,
		config:   conf,
		token:    conf.Sdk.Token,
		appId:    conf.Sdk.AppId,
		uuid:     conf.Sdk.UuId,
		validate: validate,
		tpl:      tpl,
	}
}

func (k *Khan) SetToken(token string) *Khan {
	k.token = token
	return k
}

func (k *Khan) GetToken() string {
	return k.token
}

func (k *Khan) SetAppId(appId string) *Khan {
	k.appId = appId

	return k
}

func (k *Khan) GetAppId() string {
	return k.appId
}

func (k *Khan) SetUuId(uuid string) *Khan {
	k.uuid = uuid

	return k
}
func (k *Khan) GetUuId() string {
	return k.uuid
}

func (k *Khan) SetNKey(key string) *Khan {
	k.nKey = key

	return k
}
func (k *Khan) GetNKey() string {
	return k.nKey
}

func (k *Khan) Config() *config.Config {
	return k.config
}

func (k *Khan) Client() *resty.Client {
	return k.client
}

func (k *Khan) Validate() *validator.Validate {
	return k.validate
}

func (k *Khan) Tpl() embed.FS {
	return k.tpl
}
