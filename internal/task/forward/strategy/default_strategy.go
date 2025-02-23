package strategy

import (
	"context"
	"errors"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"

	"github.com/samber/lo"
)

type DefaultStrategy struct {
	sdk *khan.Khan
}

func (ds *DefaultStrategy) Send(ctx context.Context, req v1.SnsForwardRequest, timeline v1.TimelineObject) error {
	log.C(ctx).Info().Msg("开始执行默认发送策略")
	resp, err := ds.sdk.SnsSendExcludeImage(
		ctx, sns.SnsSendExcludeImageRequest{
			AppId:        req.AppId,
			AllowUser:    lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			AllowTagId:   make([]string, 0),
			AtUser:       lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			DisableUser:  lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			DisableTagId: make([]string, 0),
			Private:      req.Privacy,
			XmlTxt:       req.SnsXml,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->snsSendExcludeImage方法失败")
		return err
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->snsSendExcludeImage方法失败")
		return errors.New("ret !=0 ->调用SnsHandler->snsSendExcludeImage方法失败")
	}

	return nil
}
