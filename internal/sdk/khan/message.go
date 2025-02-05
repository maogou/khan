package khan

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/log"
)

func (k *Khan) PostText(ctx context.Context, req transform.PostTextRequest) (*transform.PostTextResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostTextResponse{}).Post(postText)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostText方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostTextResponse), nil
}

func (k *Khan) PostImage(ctx context.Context, req transform.PostImage) (*transform.PostImageResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostImageResponse{}).Post(postImage)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostImage方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostImageResponse), nil
}

func (k *Khan) PostFile(ctx context.Context, req transform.PostFileRequest) (*transform.PostFileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostFileResponse{}).Post(sendFile)

	if err != nil {
		log.C(ctx).Error().Msg("调用postFile方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostFileResponse), nil
}

func (k *Khan) PostLink(ctx context.Context, req transform.PostLinkRequest) (*transform.PostLinkResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostLinkResponse{}).Post(sendLink)

	if err != nil {
		log.C(ctx).Error().Msg("调用postLink方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostLinkResponse), nil
}

func (k *Khan) PostVoice(ctx context.Context, req transform.PostVoiceRequest) (*transform.PostVoiceResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostVoiceResponse{}).Post(sendVoice)

	if err != nil {
		log.C(ctx).Error().Msg("调用postVoice方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostVoiceResponse), nil
}

func (k *Khan) PostVideo(ctx context.Context, req transform.PostVideoRequest) (*transform.PostVideoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostVideoResponse{}).Post(sendVideo)

	if err != nil {
		log.C(ctx).Error().Msg("调用postVideo方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostVideoResponse), nil
}

func (k *Khan) PostNameCard(ctx context.Context, req transform.PostNameCardRequest) (*transform.PostNameCardResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostNameCardResponse{}).Post(sendNameCard)

	if err != nil {
		log.C(ctx).Error().Msg("调用postNameCard方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostNameCardResponse), nil
}

func (k *Khan) PostEmoji(ctx context.Context, req transform.PostEmojiRequest) (*transform.PostEmojiResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PostEmojiResponse{}).Post(sendEmoji)

	if err != nil {
		log.C(ctx).Error().Msg("调用postEmoji方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PostEmojiResponse), nil
}

func (k *Khan) ContactList(ctx context.Context, req transform.ContactListRequest) (*transform.ContactListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactListResponse{}).Post(contactList)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactList方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactListResponse), nil
}

func (k *Khan) ContactSearch(ctx context.Context, req transform.ContactSearchRequest) (*transform.ContactSearchResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactSearchResponse{}).Post(contactSearch)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSearch方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactSearchResponse), nil
}

func (k *Khan) ContactAdd(ctx context.Context, req transform.ContactAddRequest) (*transform.ContactAddResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactAddResponse{}).Post(contactAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactAdd方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactAddResponse), nil
}

func (k *Khan) ContactDel(ctx context.Context, req transform.ContactDelRequest) (*transform.ContactDelResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactDelResponse{}).Post(contactDel)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactDel方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactDelResponse), nil
}

func (k *Khan) ContactInfo(ctx context.Context, req transform.ContactInfoRequest) (*transform.ContactInfoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactInfoResponse{}).Post(contactInfo)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactInfo方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactInfoResponse), nil
}

func (k *Khan) ContactDetail(ctx context.Context, req transform.ContactDetailRequest) (*transform.ContactDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactDetailResponse{}).Post(contactDetail)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactDetail方法失败")
		return nil, err
	}
	return resp.Result().(*transform.ContactDetailResponse), nil
}

func (k *Khan) ContactSetOnlyChat(ctx context.Context, req transform.ContactSetOnlyChatRequest) (*transform.ContactSetOnlyChatResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactSetOnlyChatResponse{}).Post(contactSetOnlyChat)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetOnlyChat方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactSetOnlyChatResponse), nil
}

func (k *Khan) ContactSetRemark(ctx context.Context, req transform.ContactSetRemarkRequest) (*transform.ContactSetRemarkResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&transform.ContactSetRemarkResponse{}).Post(contactSetRemark)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetRemark方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactSetRemarkResponse), nil
}

func (k *Khan) PersonalProfile(ctx context.Context, req transform.PersonalProfileRequest) (*transform.PersonalProfileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PersonalProfileResponse{}).Post(personalProfile)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalProfile方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalProfileResponse), nil
}

func (k *Khan) PersonalQrcode(ctx context.Context, req transform.PersonalQrcodeRequest) (*transform.PersonalQrcodResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PersonalQrcodResponse{}).Post(personalQrcode)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalQrcode方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalQrcodResponse), nil
}

func (k *Khan) PersonalSafety(ctx context.Context, req transform.PersonalSafetyRequest) (*transform.PersonalSafetyResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PersonalSafetyResponse{}).Post(personalSafety)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalSafety方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalSafetyResponse), nil
}

func (k *Khan) PersonalPrivacySetting(ctx context.Context, req transform.PersonalPrivacySettingRequest) (*transform.PersonalPrivacySettingResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PersonalPrivacySettingResponse{}).Post(personalPrivacySetting)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalPrivacySetting方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalPrivacySettingResponse), nil
}

func (k *Khan) PersonalUploadHeadImg(ctx context.Context, req transform.PersonalUploadHdHeadImgRequest) (*transform.PersonalUploadHdHeadImgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.PersonalUploadHdHeadImgResponse{}).Post(personalUploadHdHeadImg)

	if err != nil {
		log.C(ctx).Error().Msg("调用personalUploadHdHeadImg方法失败")
		return nil, err
	}

	return resp.Result().(*transform.PersonalUploadHdHeadImgResponse), nil
}

func (k *Khan) PushMsg(ctx context.Context, req v1.CollectRequest, callback string) error {
	cData := transform.CallbackRequest{
		AppId:    req.AppId,
		Data:     req.Data,
		TypeName: req.TypeName,
	}

	_, err := k.client.R().SetBody(cData).Post(callback)

	if err != nil {
		log.C(ctx).Error().Msg("回调信息PushMsg方法推送失败")
		return err
	}

	return nil
}
