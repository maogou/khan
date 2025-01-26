package gewe

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/log"
)

func (g *Gewe) PostText(ctx context.Context, req transform.PostTextRequest) (*transform.PostTextResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostTextResponse{}).Post(postText)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostText方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostTextResponse), nil
}

func (g *Gewe) PostImage(ctx context.Context, req transform.PostImage) (*transform.PostImageResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostImageResponse{}).Post(postImage)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostImage方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostImageResponse), nil
}

func (g *Gewe) PostFile(ctx context.Context, req transform.PostFileRequest) (*transform.PostFileResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostFileResponse{}).Post(sendFile)

	if err != nil {
		log.C(ctx).Error().Msg("调用postFile方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostFileResponse), nil
}

func (g *Gewe) PostLink(ctx context.Context, req transform.PostLinkRequest) (*transform.PostLinkResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostLinkResponse{}).Post(sendLink)

	if err != nil {
		log.C(ctx).Error().Msg("调用postLink方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostLinkResponse), nil
}

func (g *Gewe) PostVoice(ctx context.Context, req transform.PostVoiceRequest) (*transform.PostVoiceResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostVoiceResponse{}).Post(sendVoice)

	if err != nil {
		log.C(ctx).Error().Msg("调用postVoice方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostVoiceResponse), nil
}

func (g *Gewe) PostVideo(ctx context.Context, req transform.PostVideoRequest) (*transform.PostVideoResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostVideoResponse{}).Post(sendVideo)

	if err != nil {
		log.C(ctx).Error().Msg("调用postVideo方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostVideoResponse), nil
}

func (g *Gewe) PostNameCard(ctx context.Context, req transform.PostNameCardRequest) (*transform.PostNameCardResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostNameCardResponse{}).Post(sendNameCard)

	if err != nil {
		log.C(ctx).Error().Msg("调用postNameCard方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostNameCardResponse), nil
}

func (g *Gewe) PostEmoji(ctx context.Context, req transform.PostEmojiRequest) (*transform.PostEmojiResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PostEmojiResponse{}).Post(sendEmoji)

	if err != nil {
		log.C(ctx).Error().Msg("调用postEmoji方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostEmojiResponse), nil
}

func (g *Gewe) ContactList(ctx context.Context, req transform.ContactListRequest) (*transform.ContactListResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactListResponse{}).Post(contactList)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactList方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactListResponse), nil
}

func (g *Gewe) ContactSearch(ctx context.Context, req transform.ContactSearchRequest) (*transform.ContactSearchResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactSearchResponse{}).Post(contactSearch)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSearch方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactSearchResponse), nil
}

func (g *Gewe) PersonalProfile(ctx context.Context, req transform.PersonalProfileRequest) (*transform.PersonalProfileResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PersonalProfileResponse{}).Post(personalProfile)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalProfile方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalProfileResponse), nil
}

func (g *Gewe) PersonalQrcode(ctx context.Context, req transform.PersonalQrcodeRequest) (*transform.PersonalQrcodResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PersonalQrcodResponse{}).Post(personalQrcode)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalQrcode方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalQrcodResponse), nil
}

func (g *Gewe) PersonalSafety(ctx context.Context, req transform.PersonalSafetyRequest) (*transform.PersonalSafetyResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PersonalSafetyResponse{}).Post(personalSafety)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalSafety方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalSafetyResponse), nil
}

func (g *Gewe) PersonalPrivacySetting(ctx context.Context, req transform.PersonalPrivacySettingRequest) (*transform.PersonalPrivacySettingResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PersonalPrivacySettingResponse{}).Post(personalPrivacySetting)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalPrivacySetting方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalPrivacySettingResponse), nil
}

func (g *Gewe) PersonalUploadHeadImg(ctx context.Context, req transform.PersonalUploadHdHeadImgRequest) (*transform.PersonalUploadHdHeadImgResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.PersonalUploadHdHeadImgResponse{}).Post(personalUploadHdHeadImg)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalUploadHdHeadImg方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalUploadHdHeadImgResponse), nil
}

func (g *Gewe) PushMsg(ctx context.Context, req v1.CollectRequest, callback string) error {
	_, err := g.client.R().SetBody(req).Post(callback)

	if err != nil {
		return err
	}

	return nil
}
