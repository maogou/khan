package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/log"
)

func (k *Khan) SnsUploadImg(ctx context.Context, req sns.SnsUploadImgRequest) (*sns.SnsUploadImgResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsUploadImgResponse{}).Post(snsUploadImg)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsUploadImg方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsUploadImgResponse), nil
}

func (k *Khan) SnsMyselfPage(ctx context.Context, req sns.SnsMyselfPageRequest) (*sns.SnsMyselfPageResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsMyselfPageResponse{}).Post(snsMyselfPage)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsMyselfPage方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsMyselfPageResponse), nil
}

func (k *Khan) SnsFriendPage(ctx context.Context, req sns.SnsFriendPageRequest) (*sns.SnsFriendPageResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsFriendPageResponse{}).Post(snsFriendPage)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsFriendPage方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsFriendPageResponse), nil
}

func (k *Khan) SnsDetail(ctx context.Context, req sns.SnsDetailRequest) (*sns.SnsDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsDetailResponse{}).Post(snsDetail)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsDetail方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsDetailResponse), nil
}
