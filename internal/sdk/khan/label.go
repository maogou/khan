package khan

import (
	"context"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/log"
)

func (k *Khan) LabelAdd(ctx context.Context, req transform.LabelAddRequest) (*transform.LabelAddResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.LabelAddResponse{}).Post(labelAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelAdd方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelAddResponse), nil
}

func (k *Khan) LabelDelete(ctx context.Context, req transform.LabelDeleteRequest) (*transform.LabelDeleteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.LabelDeleteResponse{}).Post(labelDelete)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelDelete方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelDeleteResponse), nil
}

func (k *Khan) LabelList(ctx context.Context, req transform.LabelListRequest) (*transform.LabelListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.LabelListResponse{}).Post(labelList)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelList方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelListResponse), nil
}

func (k *Khan) LabelModify(ctx context.Context, req transform.LabelModifyRequest) (*transform.LabelModifyResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.LabelModifyResponse{}).Post(labelModify)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelModify方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelModifyResponse), nil
}
