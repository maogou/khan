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

func (g *Gewe) ContactAdd(ctx context.Context, req transform.ContactAddRequest) (*transform.ContactAddResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactAddResponse{}).Post(contactAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactAdd方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactAddResponse), nil
}

func (g *Gewe) ContactDel(ctx context.Context, req transform.ContactDelRequest) (*transform.ContactDelResponse, error) {

	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactDelResponse{}).Post(contactDel)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactDel方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactDelResponse), nil
}

func (g *Gewe) ContactInfo(ctx context.Context, req transform.ContactInfoRequest) (*transform.ContactInfoResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactInfoResponse{}).Post(contactInfo)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactInfo方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactInfoResponse), nil
}

func (g *Gewe) ContactDetail(ctx context.Context, req transform.ContactDetailRequest) (*transform.ContactDetailResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactDetailResponse{}).Post(contactDetail)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactDetail方法失败")
		return nil, err
	}
	return resp.Result().(*transform.ContactDetailResponse), nil
}

func (g *Gewe) ContactSetOnlyChat(ctx context.Context, req transform.ContactSetOnlyChatRequest) (*transform.ContactSetOnlyChatResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactSetOnlyChatResponse{}).Post(contactSetOnlyChat)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetOnlyChat方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactSetOnlyChatResponse), nil
}

func (g *Gewe) ContactSetRemark(ctx context.Context, req transform.ContactSetRemarkRequest) (*transform.ContactSetRemarkResponse, error) {

	resp, err := g.client.R().SetBody(req).SetResult(&transform.ContactSetRemarkResponse{}).Post(contactSetRemark)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetRemark方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ContactSetRemarkResponse), nil
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

func (g *Gewe) LabelAdd(ctx context.Context, req transform.LabelAddRequest) (*transform.LabelAddResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.LabelAddResponse{}).Post(labelAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelAdd方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelAddResponse), nil
}

func (g *Gewe) LabelDelete(ctx context.Context, req transform.LabelDeleteRequest) (*transform.LabelDeleteResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.LabelDeleteResponse{}).Post(labelDelete)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelDelete方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelDeleteResponse), nil
}

func (g *Gewe) LabelList(ctx context.Context, req transform.LabelListRequest) (*transform.LabelListResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.LabelListResponse{}).Post(labelList)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelList方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelListResponse), nil
}

func (g *Gewe) LabelModify(ctx context.Context, req transform.LabelModifyRequest) (*transform.LabelModifyResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.LabelModifyResponse{}).Post(labelModify)

	if err != nil {
		log.C(ctx).Error().Msg("调用labelModify方法失败")
		return nil, err
	}

	return resp.Result().(*transform.LabelModifyResponse), nil
}

func (g *Gewe) GroupMember(ctx context.Context, req transform.GroupMemberRequest) (*transform.GroupMemberResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupMemberResponse{}).Post(groupMember)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupMember方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupMemberResponse), nil
}

func (g *Gewe) GroupDetail(ctx context.Context, req transform.GroupDetailRequest) (*transform.GroupDetailResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupDetailResponse{}).Post(groupDetail)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupDetail方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupDetailResponse), nil
}

func (g *Gewe) GroupAnnouncement(ctx context.Context, req transform.GroupAnnouncementRequest) (*transform.GroupAnnouncementResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupAnnouncementResponse{}).Post(groupAnnouncement)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupAnnouncement方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupAnnouncementResponse), nil
}

func (g *Gewe) GroupSetAnnouncement(ctx context.Context, req transform.GroupSetAnnouncementRequest) (*transform.GroupSetAnnouncementResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupSetAnnouncementResponse{}).Post(groupSetAnnouncement)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupSetAnnouncement方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupSetAnnouncementResponse), nil
}

func (g *Gewe) GroupModifyName(ctx context.Context, req transform.GroupModifyNameRequest) (*transform.GroupModifyNameResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupModifyNameResponse{}).Post(groupModifyName)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupModifyName方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupModifyNameResponse), nil
}

func (g *Gewe) GroupModifyRemark(ctx context.Context, req transform.GroupModifyRemarkRequest) (*transform.GroupModifyRemarkResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupModifyRemarkResponse{}).Post(groupModifyRemark)
	if err != nil {
		log.C(ctx).Error().Msg("调用GroupModifyRemark方法失败")
		return nil, err
	}
	return resp.Result().(*transform.GroupModifyRemarkResponse), nil
}

func (g *Gewe) GroupModifyMyNickName(ctx context.Context, req transform.GroupModifyRemarkRequest) (*transform.GroupModifyNickNameResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupModifyNickNameResponse{}).Post(groupModifyMyNickName)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupModifyMyNickName方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupModifyNickNameResponse), nil
}

func (g *Gewe) GroupTopPing(ctx context.Context, req transform.GroupTopPingRequest) (*transform.GroupTopPingResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupTopPingResponse{}).Post(groupTopPing)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupTopPing方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupTopPingResponse), nil
}

func (g *Gewe) GroupSaveToContactList(ctx context.Context, req transform.GroupSaveContactToListRequest) (*transform.GroupSaveContactToListResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupSaveContactToListResponse{}).Post(groupContactSaveList)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupSaveToContactList方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupSaveContactToListResponse), nil
}

func (g *Gewe) GroupSetMsgSilence(ctx context.Context, req transform.GroupSetSilenceRequest) (*transform.GroupSetSilenceResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&transform.GroupSetSilenceResponse{}).Post(groupSetSilence)

	if err != nil {
		log.C(ctx).Error().Msg("调用GroupSetMsgSilence方法失败")
		return nil, err
	}

	return resp.Result().(*transform.GroupSetSilenceResponse), nil
}
func (g *Gewe) PushMsg(ctx context.Context, req v1.CollectRequest, callback string) error {
	cData := transform.CallbackRequest{
		AppId:    req.AppId,
		Data:     req.Data,
		TypeName: req.TypeName,
	}

	_, err := g.client.R().SetBody(cData).Post(callback)

	if err != nil {
		log.C(ctx).Error().Msg("回调信息PushMsg方法推送失败")
		return err
	}

	return nil
}
