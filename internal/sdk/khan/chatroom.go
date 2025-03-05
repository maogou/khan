package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/log"
)

func (k *Khan) ChatroomCreate(ctx context.Context, req chatroom.ChatroomCreateRequest) (*chatroom.ChatroomCreateResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomCreateResponse{}).Post(chatroomCreate)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomCreate方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomCreateResponse), nil
}

func (k *Khan) ChatroomMember(ctx context.Context, req chatroom.ChatroomMemberRequest) (*chatroom.ChatroomMemberResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomMemberResponse{}).Post(chatroomMember)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomMember方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomMemberResponse), nil
}

func (k *Khan) ChatroomDetail(ctx context.Context, req chatroom.ChatroomDetailRequest) (*chatroom.ChatroomDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomDetailResponse{}).Post(chatroomDetail)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomDetail方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomDetailResponse), nil
}

func (k *Khan) ChatroomAnnouncement(ctx context.Context, req chatroom.ChatroomAnnouncementRequest) (*chatroom.ChatroomAnnouncementResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomAnnouncementResponse{}).Post(chatroomAnnouncement)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomAnnouncement方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomAnnouncementResponse), nil
}

func (k *Khan) ChatroomSetAnnouncement(ctx context.Context, req chatroom.ChatroomSetAnnouncementRequest) (*chatroom.ChatroomSetAnnouncementResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomSetAnnouncementResponse{}).Post(chatroomSetAnnouncement)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomSetAnnouncement方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomSetAnnouncementResponse), nil
}

func (k *Khan) ChatroomModifyName(ctx context.Context, req chatroom.ChatroomModifyNameRequest) (*chatroom.ChatroomModifyNameResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomModifyNameResponse{}).Post(chatroomModifyName)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomModifyName方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomModifyNameResponse), nil
}

func (k *Khan) ChatroomModifyRemark(ctx context.Context, req chatroom.ChatroomModifyRemarkRequest) (*chatroom.ChatroomModifyRemarkResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomModifyRemarkResponse{}).Post(chatroomModifyRemark)
	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomModifyRemark方法失败")
		return nil, err
	}
	return resp.Result().(*chatroom.ChatroomModifyRemarkResponse), nil
}

func (k *Khan) ChatroomModifyMyNickName(ctx context.Context, req chatroom.ChatroomModifyNickNameRequest) (*chatroom.ChatroomModifyNickNameResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomModifyNickNameResponse{}).Post(chatroomModifyMyNickName)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomModifyMyNickName方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomModifyNickNameResponse), nil
}

func (k *Khan) ChatroomTopPing(ctx context.Context, req chatroom.ChatroomTopPingRequest) (*chatroom.ChatroomTopPingResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomTopPingResponse{}).Post(chatroomTopPing)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomTopPing方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomTopPingResponse), nil
}

func (k *Khan) ChatroomSaveToContactList(ctx context.Context, req chatroom.ChatroomSaveContactToListRequest) (*chatroom.ChatroomSaveContactToListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomSaveContactToListResponse{}).Post(chatroomContactSaveList)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomSaveToContactList方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomSaveContactToListResponse), nil
}

func (k *Khan) ChatroomSetMsgSilence(ctx context.Context, req chatroom.ChatroomSetSilenceRequest) (*chatroom.ChatroomSetSilenceResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomSetSilenceResponse{}).Post(chatroomSetSilence)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomSetMsgSilence方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomSetSilenceResponse), nil
}

func (k *Khan) ChatroomQrcode(ctx context.Context, req chatroom.ChatroomQrcodeRequest) (*chatroom.ChatroomQrcodeResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomQrcodeResponse{}).Post(chatroomQrcode)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomQrcode方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomQrcodeResponse), nil
}

func (k *Khan) ChatroomDeleteMember(ctx context.Context, req chatroom.ChatroomDeleteMebmerRequest) (*chatroom.ChatroomDeleteMebmerResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomDeleteMebmerResponse{}).Post(chatroomDeleteMember)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomDeleteMember方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomDeleteMebmerResponse), nil
}

func (k *Khan) ChatroomInvite(ctx context.Context, req chatroom.ChatroomInviteRequest) (*chatroom.ChatroomInviteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomInviteResponse{}).Post(chatroomInvite)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomInvite方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomInviteResponse), nil
}

func (k *Khan) ChatroomAdminOperate(ctx context.Context, req chatroom.ChatroomAdminOperateRequest) (*chatroom.ChatroomAdminOperateResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomAdminOperateResponse{}).Post(chatroomAdminOperate)

	if err != nil {
		log.C(ctx).Error().Msg("调用ChatroomAdminOperate方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomAdminOperateResponse), nil
}

func (k *Khan) AccessApplyCheckApprove(ctx context.Context, req chatroom.ChatroomConfirInviteRequest) (*chatroom.ChatroomConfirInviteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&chatroom.ChatroomConfirInviteResponse{}).Post(chatroomConfirmInvite)

	if err != nil {
		log.C(ctx).Error().Msg("调用AccessApplyCheckApprove方法失败")
		return nil, err
	}

	return resp.Result().(*chatroom.ChatroomConfirInviteResponse), nil
}
