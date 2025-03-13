package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/download"
	"smallBot/internal/pkg/log"
)

func (k *Khan) DownloadEmoji(ctx context.Context, req download.DownloadEmojiRequest) (*download.DownloadEmojiResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&download.DownloadEmojiResponse{}).Post(downloadEmoji)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用downloadEmoji方法失败")
		return nil, err
	}

	return resp.Result().(*download.DownloadEmojiResponse), nil
}

func (k *Khan) DownloadCdn(ctx context.Context, req download.DownloadCdnRequest) (*download.DownloadCdnResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&download.DownloadCdnResponse{}).Post(downloadCdn)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用downloadCdn方法失败")
		return nil, err
	}

	return resp.Result().(*download.DownloadCdnResponse), nil
}

func (k *Khan) DownloadVideo(ctx context.Context, req download.DownloadVideoRequest) (*download.DownloadVideoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&download.DownloadVideoResponse{}).Post(downloadVideo)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用downloadVideo方法失败")
		return nil, err
	}

	return resp.Result().(*download.DownloadVideoResponse), nil
}

func (k *Khan) DownloadFile(ctx context.Context, req download.DownloadFileRequest) (*download.DownloadFileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&download.DownloadFileResponse{}).Post(downloadFile)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用downloadFile方法失败")
		return nil, err
	}

	return resp.Result().(*download.DownloadFileResponse), nil
}

func (k *Khan) DownloadVoice(ctx context.Context, req download.DownloadVoiceRequest) (*download.DownloadVoiceResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&download.DownloadVoiceResponse{}).Post(downloadVoice)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用downloadVoice方法失败")
		return nil, err
	}

	return resp.Result().(*download.DownloadVoiceResponse), nil
}
