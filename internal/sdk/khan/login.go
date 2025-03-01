package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/login"
	"smallBot/internal/pkg/log"
)

func (k *Khan) Logout(ctx context.Context, req login.LogoutRequest) (*login.LogoutResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&login.LogoutResponse{}).Post(logout)

	if err != nil {
		log.C(ctx).Error().Msg("调用Logout方法失败")
		return nil, err
	}

	return resp.Result().(*login.LogoutResponse), nil
}
