package khan

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/api/khan/v1/transform/personal"
	"smallBot/internal/pkg/log"
)

func (k *Khan) PostText(ctx context.Context, req message.PostTextRequest) (*message.PostTextResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostTextResponse{}).Post(postText)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostText方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostTextResponse), nil
}

func (k *Khan) PostImage(ctx context.Context, req message.PostImage) (*message.PostImageResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostImageResponse{}).Post(postImage)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostImage方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostImageResponse), nil
}

func (k *Khan) PostFile(ctx context.Context, req message.PostFileRequest) (*message.PostFileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostFileResponse{}).Post(sendFile)

	if err != nil {
		log.C(ctx).Error().Msg("调用postFile方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostFileResponse), nil
}

func (k *Khan) PostLink(ctx context.Context, req message.PostLinkRequest) (*message.PostLinkResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostLinkResponse{}).Post(sendLink)

	if err != nil {
		log.C(ctx).Error().Msg("调用postLink方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostLinkResponse), nil
}

func (k *Khan) PostVoice(ctx context.Context, req message.PostVoiceRequest) (*message.PostVoiceResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostVoiceResponse{}).Post(sendVoice)

	if err != nil {
		log.C(ctx).Error().Msg("调用postVoice方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostVoiceResponse), nil
}

func (k *Khan) PostVideo(ctx context.Context, req message.PostVideoRequest) (*message.PostVideoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostVideoResponse{}).Post(sendVideo)

	if err != nil {
		log.C(ctx).Error().Msg("调用postVideo方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostVideoResponse), nil
}

func (k *Khan) PostNameCard(ctx context.Context, req message.PostNameCardRequest) (*message.PostNameCardResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostNameCardResponse{}).Post(sendNameCard)

	if err != nil {
		log.C(ctx).Error().Msg("调用postNameCard方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostNameCardResponse), nil
}

func (k *Khan) PostEmoji(ctx context.Context, req message.PostEmojiRequest) (*message.PostEmojiResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostEmojiResponse{}).Post(sendEmoji)

	if err != nil {
		log.C(ctx).Error().Msg("调用postEmoji方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostEmojiResponse), nil
}

func (k *Khan) ContactList(ctx context.Context, req contact.ContactListRequest) (*contact.ContactListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactListResponse{}).Post(contactList)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactList方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactListResponse), nil
}

func (k *Khan) ContactSearch(ctx context.Context, req contact.ContactSearchRequest) (*contact.ContactSearchResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactSearchResponse{}).Post(contactSearch)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSearch方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactSearchResponse), nil
}

func (k *Khan) ContactAdd(ctx context.Context, req contact.ContactAddRequest) (*contact.ContactAddResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactAddResponse{}).Post(contactAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactAdd方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactAddResponse), nil
}

func (k *Khan) ContactDel(ctx context.Context, req contact.ContactDelRequest) (*contact.ContactDelResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactDelResponse{}).Post(contactDel)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactDel方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactDelResponse), nil
}

func (k *Khan) ContactInfo(ctx context.Context, req contact.ContactInfoRequest) (*contact.ContactInfoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactInfoResponse{}).Post(contactInfo)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactInfo方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactInfoResponse), nil
}

func (k *Khan) ContactDetail(ctx context.Context, req contact.ContactDetailRequest) (*contact.ContactDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactDetailResponse{}).Post(contactDetail)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactDetail方法失败")
		return nil, err
	}
	return resp.Result().(*contact.ContactDetailResponse), nil
}

func (k *Khan) ContactSetOnlyChat(ctx context.Context, req contact.ContactSetOnlyChatRequest) (*contact.ContactSetOnlyChatResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactSetOnlyChatResponse{}).Post(contactSetOnlyChat)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetOnlyChat方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactSetOnlyChatResponse), nil
}

func (k *Khan) ContactSetRemark(ctx context.Context, req contact.ContactSetRemarkRequest) (*contact.ContactSetRemarkResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactSetRemarkResponse{}).Post(contactSetRemark)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetRemark方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactSetRemarkResponse), nil
}

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

func (k *Khan) DownloadImg(ctx context.Context, req transform.DownloadImgRequest) (*transform.DownloadImgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.DownloadImgResponse{}).Post(downloadImg)

	if err != nil {
		log.C(ctx).Error().Msg("调用downloadImg方法失败")
		return nil, err
	}

	return resp.Result().(*transform.DownloadImgResponse), nil
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
