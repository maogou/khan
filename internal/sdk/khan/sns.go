package khan

import (
	"context"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/log"
)

func (k *Khan) SnsUploadImg(ctx context.Context, req transform.SnsUploadImgRequest) (*transform.SnsUploadImgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&transform.SnsUploadImgResponse{}).Post(snsUploadImg)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsUploadImg方法失败")
		return nil, err
	}

	return resp.Result().(*transform.SnsUploadImgResponse), nil
}
