package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/label"
	"smallBot/internal/pkg/log"
)

func (k *Khan) LabelAdd(ctx context.Context, req label.LabelAddRequest) (*label.LabelAddResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&label.LabelAddResponse{}).Post(labelAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelAdd方法失败")
		return nil, err
	}

	return resp.Result().(*label.LabelAddResponse), nil
}

func (k *Khan) LabelDelete(ctx context.Context, req label.LabelDeleteRequest) (*label.LabelDeleteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&label.LabelDeleteResponse{}).Post(labelDelete)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelDelete方法失败")
		return nil, err
	}

	return resp.Result().(*label.LabelDeleteResponse), nil
}

func (k *Khan) LabelList(ctx context.Context, req label.LabelListRequest) (*label.LabelListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&label.LabelListResponse{}).Post(labelList)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelList方法失败")
		return nil, err
	}

	return resp.Result().(*label.LabelListResponse), nil
}

func (k *Khan) LabelModify(ctx context.Context, req label.LabelModifyRequest) (*label.LabelModifyResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&label.LabelModifyResponse{}).Post(labelModify)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelModify方法失败")
		return nil, err
	}

	return resp.Result().(*label.LabelModifyResponse), nil
}
