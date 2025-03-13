package setting

import (
	"github.com/gin-gonic/gin"
	"slices"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

var swithcs = map[int]string{
	1: constant.WXRedPacketCacheKey,
	2: constant.WXPaymentCacheKey,
	3: constant.WXSnsCacheKey,
}

func (s *SettingHandler) SetSwitch(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SettingHandler->SetSwitch方法")

	var (
		req     v1.AppIdsRequest
		success = make([]string, 0)
		fails   = make([]string, 0)
		useful  = make([]v1.AppIdSwitchItem, 0)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	lc, err := help.License(ctx, s.sdk.Rdb())
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取License失败")
		response.Fail(ctx, errno.GetLicenseError)
		return
	}

	appIds := lc.AppId

	for _, item := range req.AppIdSwitch {
		if slices.Contains(appIds, item.AppId) {
			useful = append(useful, item)
		} else {
			fails = append(fails, item.AppId)
		}
	}

	key := swithcs[req.Func]

	for _, item := range useful {
		cKey := key + item.AppId
		shouldValue := help.Md5(key + item.AppId)
		if *item.Switch == 0 {
			if err = s.sdk.Rdb().Del(ctx, cKey).Err(); err != nil {
				log.C(ctx).Error().Err(err).Str("appId", item.AppId).Msg("删除开关失败")
				fails = append(fails, item.AppId)
			} else {
				success = append(success, item.AppId)
			}
		} else {
			if err = s.sdk.Rdb().Set(ctx, cKey, shouldValue, 0).Err(); err != nil {
				log.C(ctx).Error().Err(err).Str("appId", item.AppId).Msg("设置开关失败")
				fails = append(fails, item.AppId)
			} else {
				success = append(success, item.AppId)
			}
		}

	}

	response.Success(
		ctx, v1.SuccessFailResponse{
			Success: success,
			Fail:    fails,
		},
	)
}
