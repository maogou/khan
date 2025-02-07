package khan

import (
	"context"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/log"
)

func (k *Khan) Sync(ctx context.Context, req transform.FavorSyncRequest) (*transform.FavorSyncResponse, error) {
	resp, err := k.client.R().SetBody(req).SetBody(&transform.FavorSyncResponse{}).Post(favorSync)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorSync方法失败")
		return nil, err
	}

	return resp.Result().(*transform.FavorSyncResponse), nil
}

func (k *Khan) Detail(ctx context.Context, req transform.FavorDetailRequest) (*transform.FavorDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.FavorDetailResponse{}).Post(favorDetail)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorDetail方法失败")
		return nil, err
	}

	return resp.Result().(*transform.FavorDetailResponse), nil
}

func (k *Khan) Delete(ctx context.Context, req transform.FavorDeleteRequest) (*transform.FavorDeleteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.FavorDeleteResponse{}).Post(favorDelete)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorDelete方法失败")
		return nil, err
	}

	return resp.Result().(*transform.FavorDeleteResponse), nil
}
