package help

import (
	"context"
	"encoding/json"
	"errors"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/log"
	"strings"

	"github.com/redis/go-redis/v9"
)

func Permission(ctx context.Context, rdb *redis.Client) (*v1.Permission, error) {
	var p v1.Permission

	lic, err := License(ctx, rdb)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(lic.Dat, &p); err != nil {
		return nil, err
	}

	if lic.Expired() {
		return nil, errors.New("授权许可证已过期")
	}

	return &p, nil
}

func License(ctx context.Context, rdb *redis.Client) (*license.License, error) {
	var pKey string

	pKey, err := rdb.Get(ctx, constant.LicenseKey).Result()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("Redis 获取授权许可证key为空")

		return nil, err
	}

	lb, err := rdb.Get(ctx, constant.License).Result()

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("Redis 获取授权许可证为空")
		return nil, err
	}

	pKey = strings.NewReplacer(
		constant.License37, "+",
		constant.License73, "/",
		constant.License919, "=",
	).Replace(pKey)

	lic, err := license.Parse(pKey, lb)

	return lic, err
}
