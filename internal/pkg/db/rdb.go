package db

import (
	"context"
	"maogou/khan/internal/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func MustInitRedis(conf *config.Config) *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		},
	)

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()

	if err != nil {
		log.Error().Err(err).Msg("redis连接失败")
		panic(err)
	}

	return client

}
