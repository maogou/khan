package strategy

import (
	"context"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"

	"github.com/samber/lo"
)

type DefaultStrategy struct {
	sdk *khan.Khan
}

func (ds *DefaultStrategy) Send(ctx context.Context, req v1.SnsForwardRequest, timeline v1.TimelineObject) error {
	log.C(ctx).Info().Msg("开始执行默认发送策略")
	resp, err := ds.sdk.SnsSendExcludeImage(
		ctx, v1.SnsForwardRequest{
			AppId:        req.AppId,
			AllowWxIds:   lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			AtWxIds:      lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			DisableWxIds: lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			DisableTags:  make([]string, 0),
			Privacy:      req.Privacy,
			SnsXml:       req.SnsXml,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->snsSendExcludeImage方法失败")
		return err
	}

	if resp.Ret != 200 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->snsSendExcludeImage方法失败")
		return errors.New("ret !=0 ->调用SnsHandler->snsSendExcludeImage方法失败")
	}

	return nil
}
