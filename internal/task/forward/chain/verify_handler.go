package chain

import (
	"context"
	"encoding/json"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/license"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"slices"
	"strings"
)

type VerifyHandler struct {
	BaseHandler
	sdk *khan.Khan
}

var _ Handler = (*VerifyHandler)(nil)

func NewVerifyHandler(sdk *khan.Khan) Handler {
	return &VerifyHandler{
		sdk: sdk,
	}
}

func (v *VerifyHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行许可证是否过期验证.....")

	var (
		p    v1.Permission
		pKey string
	)

	keys := []string{constant.License, constant.LicenseKey}
	vals, err := v.sdk.Rdb().MGet(ctx, keys...).Result()

	lb, ok := vals[0].(string)
	if !ok || len(lb) == 0 {
		log.C(ctx).Error().Str("appid", pld.AppID).Msg("Redis 获取授权许可证为空")
		return errors.New("Redis 获取授权许可证为空")
	}

	pKey, ok = vals[1].(string)
	if !ok || len(pKey) == 0 {
		log.C(ctx).Error().Str("appid", pld.AppID).Msg("Redis 获取授权许可证key为空")

		return errors.New("Redis 获取授权许可证解密key为空")
	}

	pKey = strings.NewReplacer(
		constant.License37, "+",
		constant.License73, "/",
		constant.License919, "=",
	).Replace(pKey)

	lic, err := license.Parse(pKey, lb)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("license解析失败")
		return errors.New("license解析失败")
	}

	if err = json.Unmarshal(lic.Dat, &p); err != nil {
		log.C(ctx).Error().Err(err).Msg("json解析授权许可证失败")
		return errors.New("json解析授权许可证失败")
	}

	if lic.Expired() {
		log.C(ctx).Error().Msg("许可证已过期")
		return errors.New("许可证已过期")
	}

	log.C(ctx).Info().Any("p", p).Msg("用户申请的权限")

	if !slices.Contains(p.AppId, pld.AppID) {
		log.C(ctx).Warn().Msg("对应的wxid未授权，无法使用朋友圈转发服务")
		return errors.New("对应的wxid未授权，无法使用朋友圈转发服务")
	}

	if sns, sOk := p.Permission[constant.SnsFun]; !sOk || sns == 0 {
		log.C(ctx).Warn().Msg("对应的wxid未授权，没有开通朋友圈功能,无法使用朋友圈转发服务")
		return errors.New("对应的wxid未授权，没有开通朋友圈功能,无法使用朋友圈转发服务")
	}

	return nil
}
