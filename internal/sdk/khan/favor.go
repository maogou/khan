package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/favor"
	"smallBot/internal/pkg/log"
)

func (k *Khan) Sync(ctx context.Context, req favor.FavorSyncRequest) (*favor.FavorSyncResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&favor.FavorSyncResponse{}).Post(favorSync)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorSync方法失败")
		return nil, err
	}

	return resp.Result().(*favor.FavorSyncResponse), nil
}

func (k *Khan) Detail(ctx context.Context, req favor.FavorDetailRequest) (*favor.FavorDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&favor.FavorDetailResponse{}).Post(favorDetail)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorDetail方法失败")
		return nil, err
	}

	return resp.Result().(*favor.FavorDetailResponse), nil
}

func (k *Khan) Delete(ctx context.Context, req favor.FavorDeleteRequest) (*favor.FavorDeleteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&favor.FavorDeleteResponse{}).Post(favorDelete)

	if err != nil {
		log.C(ctx).Error().Msg("调用favorDelete方法失败")
		return nil, err
	}

	return resp.Result().(*favor.FavorDeleteResponse), nil
}
