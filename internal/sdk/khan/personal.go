package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/personal"
	"smallBot/internal/pkg/log"
)

func (k *Khan) PersonalProfile(ctx context.Context, req personal.PersonalProfileRequest) (*personal.PersonalProfileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&personal.PersonalProfileResponse{}).Post(personalProfile)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalProfile方法失败")
		return nil, err
	}

	return resp.Result().(*personal.PersonalProfileResponse), nil
}

func (k *Khan) PersonalQrcode(ctx context.Context, req personal.PersonalQrcodeRequest) (*personal.PersonalQrcodResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&personal.PersonalQrcodResponse{}).Post(personalQrcode)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalQrcode方法失败")
		return nil, err
	}

	return resp.Result().(*personal.PersonalQrcodResponse), nil
}

func (k *Khan) PersonalSafety(ctx context.Context, req personal.PersonalSafetyRequest) (*personal.PersonalSafetyResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&personal.PersonalSafetyResponse{}).Post(personalSafety)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalSafety方法失败")
		return nil, err
	}

	return resp.Result().(*personal.PersonalSafetyResponse), nil
}

func (k *Khan) PersonalPrivacySetting(ctx context.Context, req personal.PersonalPrivacySettingRequest) (*personal.PersonalPrivacySettingResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&personal.PersonalPrivacySettingResponse{}).Post(personalPrivacySetting)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalPrivacySetting方法失败")
		return nil, err
	}

	return resp.Result().(*personal.PersonalPrivacySettingResponse), nil
}

func (k *Khan) PersonalUploadHeadImg(ctx context.Context, req personal.PersonalUploadHdHeadImgRequest) (*personal.PersonalUploadHdHeadImgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&personal.PersonalUploadHdHeadImgResponse{}).Post(personalUploadHdHeadImg)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalUploadHdHeadImg方法失败")
		return nil, err
	}

	return resp.Result().(*personal.PersonalUploadHdHeadImgResponse), nil
}

func (k *Khan) PersonalUpdateProfile(ctx context.Context, req personal.PersonalUpdateProfileRequest) (*personal.PersonalUpdateProfileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&personal.PersonalUpdateProfileResponse{}).Post(personalUpdateProfile)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalUpdateProfile方法失败")
		return nil, err
	}

	return resp.Result().(*personal.PersonalUpdateProfileResponse), nil
}
