package khan

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/download"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/log"
)

func (k *Khan) PostText(ctx context.Context, req v1.PostTextRequest) (*v1.PostTextResponse, error) {
	resp, err := k.tRequest().SetBody(req).SetResult(&v1.PostTextResponse{}).Post("/api/message/postText")
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostText方法失败")
		return nil, err
	}

	return resp.Result().(*v1.PostTextResponse), nil
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

func (k *Khan) DownloadImg(ctx context.Context, req download.DownloadImgRequest) (*download.DownloadImgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&download.DownloadImgResponse{}).Post(downloadImg)

	if err != nil {
		log.C(ctx).Error().Msg("调用downloadImg方法失败")
		return nil, err
	}

	return resp.Result().(*download.DownloadImgResponse), nil
}

func (k *Khan) PushMsg(ctx context.Context, req v1.CollectRequest, callback, wxid string) error {
	cData := message.CallbackRequest{
		AppId:    req.AppId,
		Data:     req.Data,
		TypeName: req.TypeName,
		Wxid:     wxid,
	}

	log.C(ctx).Info().Any("push_data", cData).Msg("推送的消息")

	cKey := constant.WXCallbackCache + req.AppId

	if customCallback, err := k.rdb.Get(ctx, cKey).Result(); err != nil {
		log.C(ctx).Error().Err(err).Msg("获取自定义回调地址失败")
	} else if customCallback != "" {
		callback = customCallback
	}

	log.C(ctx).Info().Str("callback", callback).Msg("回调地址")

	_, err := k.client.R().SetBody(cData).Post(callback)

	if err != nil {
		log.C(ctx).Error().Msg("回调信息PushMsg方法推送失败")
		return err
	}

	return nil
}

func (k *Khan) RevokeMsg(ctx context.Context, req message.RevokeMsgRequest) (*message.RevokeMsgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.RevokeMsgResponse{}).Post(revokeMsg)

	if err != nil {
		log.C(ctx).Error().Msg("调用revokeMsg方法失败")
		return nil, err
	}

	return resp.Result().(*message.RevokeMsgResponse), nil
}

func (k *Khan) ForwardMiniApp(ctx context.Context, req message.ForwardMiniAppRequest) (*message.ForwardMiniAppResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.ForwardMiniAppResponse{}).Post(sendApplet)

	if err != nil {
		log.C(ctx).Error().Msg("调用forwardMiniApp方法失败")
		return nil, err
	}

	return resp.Result().(*message.ForwardMiniAppResponse), nil
}

func (k *Khan) ForwardUrl(ctx context.Context, req message.ForwardUrlRequest) (*message.ForwardUrlResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.ForwardUrlResponse{}).Post(forwardUrl)

	if err != nil {
		log.C(ctx).Error().Msg("调用forwardUrl方法失败")
		return nil, err
	}

	return resp.Result().(*message.ForwardUrlResponse), nil
}

func (k *Khan) ForwardVideo(ctx context.Context, req message.ForwardVideoRequest) (*message.ForwardVideoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.ForwardVideoResponse{}).Post(forwardVideo)

	if err != nil {
		log.C(ctx).Error().Msg("调用forwardVideo方法失败")
		return nil, err
	}

	return resp.Result().(*message.ForwardVideoResponse), nil
}

func (k *Khan) ForwardImage(ctx context.Context, req message.ForwardImageRequest) (*message.ForwardImageResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.ForwardImageResponse{}).Post(forwardImage)

	if err != nil {
		log.C(ctx).Error().Msg("调用forwardImage方法失败")
		return nil, err
	}

	return resp.Result().(*message.ForwardImageResponse), nil
}

func (k *Khan) ForwardFile(ctx context.Context, req message.ForwardFileRequest) (*message.ForwardFileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.ForwardFileResponse{}).Post(forwardFile)

	if err != nil {
		log.C(ctx).Error().Msg("调用forwardFile方法失败")
		return nil, err
	}

	return resp.Result().(*message.ForwardFileResponse), nil
}

func (k *Khan) SendMiniApp(ctx context.Context, req message.SendMiniAppRequest) (*message.SendMiniAppResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.SendMiniAppResponse{}).Post(sendApplet)

	if err != nil {
		log.C(ctx).Error().Msg("调用sendMiniApp方法失败")
		return nil, err
	}

	return resp.Result().(*message.SendMiniAppResponse), nil
}

func (k *Khan) PostAppMsg(ctx context.Context, req message.PostAppMsgRequest) (*message.PostAppMsgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&message.PostAppMsgResponse{}).Post(postAppMsg)

	if err != nil {
		log.C(ctx).Error().Msg("调用postAppMsg方法失败")
		return nil, err
	}

	return resp.Result().(*message.PostAppMsgResponse), nil
}
