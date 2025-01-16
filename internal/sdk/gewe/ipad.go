package gewe

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/help"
)

func (g *Gewe) LoginQrCode(ctx context.Context, req v1.LoginQrCodeRequest) (*v1.LoginQrCodeResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用LoginQrCode方法")

	resp, err := g.client.R().SetBody(req).SetResult(&v1.LoginQrCodeResponse{}).Post(loginQrCode)

	if err != nil {
		log.Error().Msg("调用LoginQrCode方法失败")
		return nil, err
	}

	return resp.Result().(*v1.LoginQrCodeResponse), nil
}

func (g *Gewe) CheckLoginQrCode(ctx context.Context, req v1.CheckLoginQrCodeRequest) (*v1.CheckLoginQrCodeResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用CheckLoginQrCode方法")

	resp, err := g.client.R().SetBody(req).SetResult(&v1.CheckLoginQrCodeResponse{}).Post(checkLoginQrCode)

	if err != nil {
		log.Error().Msg("调用CheckLoginQrCode方法失败")
		return nil, err
	}

	return resp.Result().(*v1.CheckLoginQrCodeResponse), nil
}

func (g *Gewe) LongOpen(ctx context.Context, req v1.LongOpenRequest) (*v1.LongOpenResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用LongOpen方法")

	api := g.config.Sdk.Long + longOpen
	resp, err := g.client.R().SetBody(req).SetResult(&v1.LongOpenResponse{}).Post(api)

	if err != nil {
		log.Error().Msg("调用LongOpen方法失败")
		return nil, err
	}

	return resp.Result().(*v1.LongOpenResponse), nil
}

func (g *Gewe) SecAutoAuth(ctx context.Context, req v1.SecAutoAuthRequest) (*v1.SecAutoAuthResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用SecAutoAuth方法")

	resp, err := g.client.R().SetBody(req).SetResult(&v1.SecAutoAuthResponse{}).Post(secAutoAuth)

	if err != nil {
		log.Error().Msg("调用SecAutoAuth方法失败")
		return nil, err
	}

	return resp.Result().(*v1.SecAutoAuthResponse), nil
}

func (g *Gewe) CreateApp(ctx context.Context, req v1.CreateAppRequest) (*v1.CreateAppResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用CreateApp方法")

	resp, err := g.client.R().SetBody(req).SetResult(&v1.CreateAppResponse{}).Post(createApp)

	if err != nil {
		log.Error().Msg("调用CreateApp方法失败")
		return nil, err
	}

	return resp.Result().(*v1.CreateAppResponse), nil
}

func (g *Gewe) HearBeat(ctx context.Context, req v1.HearBeatRequest) (*v1.HearBeatResponse, error) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用HearBeat方法")

	resp, err := g.client.R().SetBody(req).SetResult(&v1.HearBeatResponse{}).Post(heartbeat)

	if err != nil {
		log.Error().Msg("调用HearBeat方法失败")
		return nil, err
	}

	return resp.Result().(*v1.HearBeatResponse), nil
}
