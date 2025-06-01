package khan

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
)

func (k *Khan) Sync(ctx context.Context, req v1.FavorSyncRequest) (*v1.FavorSyncResponse, error) {
	resp, err := k.tRequest().SetBody(req).SetResult(&v1.FavorSyncResponse{}).Post("/api/favor/sync")

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用favorSync方法失败")
		return nil, err
	}

	return resp.Result().(*v1.FavorSyncResponse), nil
}

func (k *Khan) Detail(ctx context.Context, req v1.FavorDetailRequest) (*v1.FavorDetailResponse, error) {
	resp, err := k.tRequest().SetBody(req).SetResult(&v1.FavorDetailResponse{}).Post("/api/favor/getContent")

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用favorDetail方法失败")
		return nil, err
	}

	return resp.Result().(*v1.FavorDetailResponse), nil
}

func (k *Khan) Delete(ctx context.Context, req v1.FavorDeleteRequest) (*v1.ResponseSuccessWithoutData, error) {
	resp, err := k.tRequest().SetBody(req).SetResult(&v1.ResponseSuccessWithoutData{}).Post("/api/favor/delete")

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用favorDelete方法失败")
		return nil, err
	}

	return resp.Result().(*v1.ResponseSuccessWithoutData), nil
}
