package help

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"strings"
)

func License(ctx context.Context, rdb *redis.Client) (*v1.Permission, error) {
	var (
		p    v1.Permission
		pKey string
	)
	keys := []string{constant.License, constant.LicenseKey}

	vals, err := rdb.MGet(ctx, keys...).Result()

	lb, ok := vals[0].(string)
	if !ok || len(lb) == 0 {
		return nil, errors.New("Redis 获取授权许可证解密key为空")
	}

	pKey, ok = vals[1].(string)
	if !ok || len(pKey) == 0 {
		return nil, errors.New("Redis 获取授权许可证解密key为空")
	}

	pKey = strings.NewReplacer(
		constant.License37, "+",
		constant.License73, "/",
		constant.License919, "=",
	).Replace(pKey)

	lic, err := license.Parse(pKey, lb)

	if err = json.Unmarshal(lic.Dat, &p); err != nil {
		return nil, err
	}

	if lic.Expired() {
		return nil, errors.New("授权许可证已过期")
	}

	return &p, nil
}
