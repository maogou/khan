package khan

import (
	"smallBot/internal/config"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
)

type Khan struct {
	client   *resty.Client
	config   *config.Config
	appId    string
	uuid     string
	nKey     string
	validate *validator.Validate
	rdb      *redis.Client
}

func NewKhanSdk(conf *config.Config, client *resty.Client, validate *validator.Validate, rdb *redis.Client) *Khan {
	client.SetTimeout(conf.Sdk.TimeOut * time.Second)
	client.BaseURL = conf.Sdk.Gog7a6v8g

	return &Khan{
		client:   client,
		config:   conf,
		validate: validate,
		rdb:      rdb,
	}
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

func (k *Khan) Rdb() *redis.Client {
	return k.rdb
}
