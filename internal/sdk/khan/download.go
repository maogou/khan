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
