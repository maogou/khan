package khan

import (
	"context"
	"maogou/khan/api/khan/v1/transform/login"
	"maogou/khan/internal/pkg/log"
)

func (k *Khan) Logout(ctx context.Context, req login.LogoutRequest) (*login.LogoutResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&login.LogoutResponse{}).Post(logout)

	if err != nil {
		log.C(ctx).Error().Msg("调用Logout方法失败")
		return nil, err
	}

	return resp.Result().(*login.LogoutResponse), nil
}
