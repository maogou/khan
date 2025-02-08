package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/label"
	"smallBot/internal/pkg/log"
)

func (k *Khan) Sync(ctx context.Context, req label.FavorSyncRequest) (*label.FavorSyncResponse, error) {
	resp, err := k.client.R().SetBody(req).SetBody(&label.FavorSyncResponse{}).Post(favorSync)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorSync方法失败")
		return nil, err
	}

	return resp.Result().(*label.FavorSyncResponse), nil
}

func (k *Khan) Detail(ctx context.Context, req label.FavorDetailRequest) (*label.FavorDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&label.FavorDetailResponse{}).Post(favorDetail)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorDetail方法失败")
		return nil, err
	}

	return resp.Result().(*label.FavorDetailResponse), nil
}

func (k *Khan) Delete(ctx context.Context, req label.FavorDeleteRequest) (*label.FavorDeleteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&label.FavorDeleteResponse{}).Post(favorDelete)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorDelete方法失败")
		return nil, err
	}

	return resp.Result().(*label.FavorDeleteResponse), nil
}
