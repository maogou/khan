package khan

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/login"
	"smallBot/internal/pkg/help"
)

func (k *Khan) LoginQrCode(ctx context.Context, req login.GetLoginQrCodeRequest) (*login.GetLoginQrCodeResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用LoginQrCode方法")

	resp, err := k.client.R().SetBody(req).SetResult(&login.GetLoginQrCodeResponse{}).Post(loginQrCode)

	if err != nil {
		log.Error().Msg("调用LoginQrCode方法失败")
		return nil, err
	}

	return resp.Result().(*login.GetLoginQrCodeResponse), nil
}

func (k *Khan) CheckLoginQrCode(ctx context.Context, req login.CheckLoginRequest) (*login.CheckLoginResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用CheckLoginQrCode方法")

	resp, err := k.client.R().SetBody(req).SetResult(&login.CheckLoginResponse{}).Post(checkLoginQrCode)

	if err != nil {
		log.Error().Msg("调用CheckLoginQrCode方法失败")
		return nil, err
	}

	return resp.Result().(*login.CheckLoginResponse), nil
}

func (k *Khan) LongOpen(ctx context.Context, req v1.LongOpenRequest) (*v1.LongOpenResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用LongOpen方法")

	api := k.config.Sdk.Long + longOpen
	resp, err := k.client.R().SetBody(req).SetResult(&v1.LongOpenResponse{}).Post(api)

	if err != nil {
		log.Error().Msg("调用LongOpen方法失败")
		return nil, err
	}

	return resp.Result().(*v1.LongOpenResponse), nil
}

func (k *Khan) SecAutoAuth(ctx context.Context, req v1.SecAutoAuthRequest) (*v1.SecAutoAuthResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用SecAutoAuth方法")

	resp, err := k.client.R().SetBody(req).SetResult(&v1.SecAutoAuthResponse{}).Post(secAutoAuth)

	if err != nil {
		log.Error().Msg("调用SecAutoAuth方法失败")
		return nil, err
	}

	return resp.Result().(*v1.SecAutoAuthResponse), nil
}

func (k *Khan) CreateApp(ctx context.Context, req v1.CreateAppRequest) (*v1.CreateAppResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用CreateApp方法")

	resp, err := k.client.R().SetBody(req).SetResult(&v1.CreateAppResponse{}).Post(createApp)

	if err != nil {
		log.Error().Msg("调用CreateApp方法失败")
		return nil, err
	}

	return resp.Result().(*v1.CreateAppResponse), nil
}

func (k *Khan) HearBeat(ctx context.Context, req v1.HearBeatRequest) (*v1.HearBeatResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用HearBeat方法")

	resp, err := k.client.R().SetBody(req).SetResult(&v1.HearBeatResponse{}).Post(heartbeat)

	if err != nil {
		log.Error().Msg("调用HearBeat方法失败")
		return nil, err
	}

	return resp.Result().(*v1.HearBeatResponse), nil
}
