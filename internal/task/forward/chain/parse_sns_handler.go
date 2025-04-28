package chain

import (
	"context"
	"encoding/xml"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
)

type ParseSnsHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewPareSnsHandler(sdk *khan.Khan) Handler {
	return &ParseSnsHandler{
		sdk: sdk,
	}
}

var _ Handler = (*ParseSnsHandler)(nil)

func (p *ParseSnsHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("snsId", pld.SnsID).Msg("开始执行解析sns.....")

	var (
		timeline v1.TimelineObject
		req      v1.SnsForwardRequest
	)

	if err := xml.Unmarshal(
		[]byte(pld.SnsDetail.Data.Object.ObjectDesc.Buffer), &timeline,
	); err != nil {
		log.C(ctx).Error().Err(err).Msg("SnsDetail->xml反序列化失败")
		return errors.New("SnsDetail->xml反序列化失败")
	}

	pld.ContentStyle = timeline.ContentObject.ContentStyle
	req = v1.SnsForwardRequest{
		AppId:   pld.AppID,
		Privacy: false,
		DisableWxIds: []string{
			pld.SnsDetail.Data.Object.Username,
		},
		SnsXml: pld.SnsDetail.Data.Object.ObjectDesc.Buffer,
	}

	pld.Req = req
	pld.Timeline = timeline

	return nil
}
