package chain

import (
	"context"
	"encoding/xml"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"strings"
)

type ParseFavHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewParseFavHandler(sdk *khan.Khan) Handler {
	return &ParseFavHandler{
		sdk: sdk,
	}
}

var _ Handler = (*ParseFavHandler)(nil)

func (pf *ParseFavHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行解析收藏夹,获取朋友圈id.....")

	if len(pld.FavDetail.Content) == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中内容为空,无需转发")
		return errors.New("收藏夹中内容为空,无需转发")
	}

	favContent := pld.FavDetail.Content

	var item v1.FavItem
	if xErr := xml.Unmarshal([]byte(favContent), &item); xErr != nil {
		log.C(ctx).Error().Err(xErr).Msg("收藏夹中内容,xml反序列化失败")
		return errors.New("收藏夹中内容,xml反序列化失败")
	}

	if len(item.Source.SourceID) == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中内容,获取不到朋友圈id,无需转发")
		return errors.New("收藏夹中内容,获取不到朋友圈id,无需转发")
	}

	log.C(ctx).Info().Str("source_sns_id===>", item.Source.SourceID).Msg("需要解析的原始snsid")

	sourceSlice := strings.Split(item.Source.SourceID, "#")

	if len(sourceSlice) == 1 {
		sourceSlice = strings.Split(item.Source.SourceID, "_")
	}

	log.C(ctx).Info().Any("sourceSlice", sourceSlice).Msg("使用#或者_分割后的朋友圈id")

	if len(sourceSlice) > 1 {
		pld.SnsID = sourceSlice[0]
	}

	return nil
}
