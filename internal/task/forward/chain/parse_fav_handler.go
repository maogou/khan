package chain

import (
	"context"
	"encoding/xml"
	"errors"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
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

	if len(pld.FavDetail.Data.ObjectList) == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中内容为空,无需转发")
		return errors.New("收藏夹中内容为空,无需转发")
	}

	favContent := pld.FavDetail.Data.ObjectList[0].Object

	if len(favContent) == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中内容为空,无需转发")
		return errors.New("收藏夹中内容为空,无需转发")
	}

	var item v1.FavItem
	if xErr := xml.Unmarshal([]byte(favContent), &item); xErr != nil {
		log.C(ctx).Error().Err(xErr).Msg("收藏夹中内容,xml反序列化失败")
		return errors.New("收藏夹中内容,xml反序列化失败")
	}

	if len(item.Source.SourceID) == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中内容,获取不到朋友圈id,无需转发")
		return errors.New("收藏夹中内容,获取不到朋友圈id,无需转发")
	}

	sourceSlice := strings.Split(item.Source.SourceID, "#")

	if len(sourceSlice) > 1 {
		pld.SnsID = sourceSlice[0]
	}

	return nil
}
