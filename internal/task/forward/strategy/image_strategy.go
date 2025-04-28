package strategy

import (
	"context"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"strconv"

	"github.com/samber/lo"
	"github.com/spf13/cast"
)

type ImageStrategy struct {
	sdk *khan.Khan
}

func (is *ImageStrategy) Send(ctx context.Context, req v1.SnsForwardRequest, timeline v1.TimelineObject) error {
	log.C(ctx).Info().Msg("开始执行图片发送策略")
	var transformImage []sns.SnsSendImageItem

	for _, media := range timeline.ContentObject.MediaList.Media {
		transformImage = append(
			transformImage, sns.SnsSendImageItem{
				ThumbUrl:  media.Thumb.Text,
				Width:     strconv.Itoa(cast.ToInt(media.Size.Width)),
				TotalSize: media.Size.TotalSize,
				Id:        media.Id,
				Type:      "1",
				Url:       media.Url.Text,
				Height:    strconv.Itoa(cast.ToInt(media.Size.Height)),
				Md5:       media.Url.Md5,
			},
		)
	}

	resp, err := is.sdk.SnsSendImage(
		ctx, sns.SnsSendImageRequest{
			AppId:        req.AppId,
			AllowUser:    lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			AllowTagId:   make([]string, 0),
			AtUser:       lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			DisableUser:  lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			DisableTagId: make([]string, 0),
			Private:      req.Privacy,
			XmlTxt:       req.SnsXml,
			Content:      timeline.ContentDesc,
			Media:        transformImage,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->snsSendImage方法失败")
		return err
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->snsSendImage方法失败")
		return errors.New("ret !=0 ->调用SnsHandler->snsSendImage方法失败")
	}

	return nil
}
