package khan

import (
	"context"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/log"
)

func (k *Khan) ChatroomCreate(ctx context.Context, req transform.ChatroomCreateRequest) (*transform.ChatroomCreateResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomCreateResponse{}).Post(chatroomCreate)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomCreate方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomCreateResponse), nil
}

func (k *Khan) ChatroomMember(ctx context.Context, req transform.ChatroomMemberRequest) (*transform.ChatroomMemberResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomMemberResponse{}).Post(chatroomMember)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomMember方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomMemberResponse), nil
}

func (k *Khan) ChatroomDetail(ctx context.Context, req transform.ChatroomDetailRequest) (*transform.ChatroomDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomDetailResponse{}).Post(chatroomDetail)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomDetail方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomDetailResponse), nil
}

func (k *Khan) ChatroomAnnouncement(ctx context.Context, req transform.ChatroomAnnouncementRequest) (*transform.ChatroomAnnouncementResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomAnnouncementResponse{}).Post(chatroomAnnouncement)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomAnnouncement方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomAnnouncementResponse), nil
}

func (k *Khan) ChatroomSetAnnouncement(ctx context.Context, req transform.ChatroomSetAnnouncementRequest) (*transform.ChatroomSetAnnouncementResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomSetAnnouncementResponse{}).Post(chatroomSetAnnouncement)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomSetAnnouncement方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomSetAnnouncementResponse), nil
}

func (k *Khan) ChatroomModifyName(ctx context.Context, req transform.ChatroomModifyNameRequest) (*transform.ChatroomModifyNameResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomModifyNameResponse{}).Post(chatroomModifyName)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomModifyName方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomModifyNameResponse), nil
}

func (k *Khan) ChatroomModifyRemark(ctx context.Context, req transform.ChatroomModifyRemarkRequest) (*transform.ChatroomModifyRemarkResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomModifyRemarkResponse{}).Post(chatroomModifyRemark)
	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomModifyRemark方法失败")
		return nil, err
	}
	return resp.Result().(*transform.ChatroomModifyRemarkResponse), nil
}

func (k *Khan) ChatroomModifyMyNickName(ctx context.Context, req transform.ChatroomModifyNickNameRequest) (*transform.ChatroomModifyNickNameResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomModifyNickNameResponse{}).Post(chatroomModifyMyNickName)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomModifyMyNickName方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomModifyNickNameResponse), nil
}

func (k *Khan) ChatroomTopPing(ctx context.Context, req transform.ChatroomTopPingRequest) (*transform.ChatroomTopPingResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomTopPingResponse{}).Post(chatroomTopPing)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomTopPing方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomTopPingResponse), nil
}

func (k *Khan) ChatroomSaveToContactList(ctx context.Context, req transform.ChatroomSaveContactToListRequest) (*transform.ChatroomSaveContactToListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomSaveContactToListResponse{}).Post(chatroomContactSaveList)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomSaveToContactList方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomSaveContactToListResponse), nil
}

func (k *Khan) ChatroomSetMsgSilence(ctx context.Context, req transform.ChatroomSetSilenceRequest) (*transform.ChatroomSetSilenceResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.ChatroomSetSilenceResponse{}).Post(chatroomSetSilence)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomSetMsgSilence方法失败")
		return nil, err
	}

	return resp.Result().(*transform.ChatroomSetSilenceResponse), nil
}
