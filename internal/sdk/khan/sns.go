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

func (k *Khan) SnsLike(ctx context.Context, req sns.SnsLikeRequest) (*sns.SnsLikeResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsLikeResponse{}).Post(snsLike)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsLike方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsLikeResponse), nil
}

func (k *Khan) SnsCancelLike(ctx context.Context, req sns.SnsCancelLikeRequest) (*sns.SnsCancelLikeResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsCancelLikeResponse{}).Post(snsCancelLike)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsCancelLike方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsCancelLikeResponse), nil
}

func (k *Khan) SnsComment(ctx context.Context, req sns.SnsCommentRequest) (*sns.SnsCommentResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsCommentResponse{}).Post(snsComment)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsComment方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsCommentResponse), nil
}

func (k *Khan) SnsDeleteComment(ctx context.Context, req sns.SnsDeleteCommentRequest) (*sns.SnsDeleteCommentResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsDeleteCommentResponse{}).Post(snsDeleteComment)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsDeleteComment方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsDeleteCommentResponse), nil
}

func (k *Khan) SnsSetPrivacyScope(ctx context.Context, req sns.SnsSetPrivacyScopeRequest) (*sns.SnsSetPrivacyScopeResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsSetPrivacyScopeResponse{}).Post(snsSetPrivacyScope)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsSetPrivacyScope方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsSetPrivacyScopeResponse), nil
}

func (k *Khan) SnsStrangerVisibilityEnabled(ctx context.Context, req sns.SnsStrangerVisibilityEnabledRequest) (*sns.SnsStrangerVisibilityEnabledResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsStrangerVisibilityEnabledResponse{}).Post(snsStrangerVisibilityEnabled)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsStrangerVisibilityEnabled方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsStrangerVisibilityEnabledResponse), nil
}

func (k *Khan) SnsSetPrivacy(ctx context.Context, req sns.SnsSetPrivacyRequest) (*sns.SnsSetPrivacyReponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsSetPrivacyReponse{}).Post(snsSetPrivacy)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsSetPrivacy方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsSetPrivacyReponse), nil
}

func (k *Khan) SnsDelete(ctx context.Context, req sns.SnsDeleteRequest) (*sns.SnsDeleteResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsDeleteResponse{}).Post(snsDelete)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsDelete方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsDeleteResponse), nil
}

func (k *Khan) SnsUploadVideo(ctx context.Context, req sns.SnsUploadVideoRequest) (*sns.SnsUploadVideoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsUploadVideoResponse{}).Post(snsUploadVideo)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsUploadVideo方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsUploadVideoResponse), nil
}

func (k *Khan) SnsSendText(ctx context.Context, req sns.SnsSendTextRequest) (*sns.SnsSendTextResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsSendTextResponse{}).Post(snsSendText)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsSendText方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsSendTextResponse), nil
}

func (k *Khan) SnsSendImage(ctx context.Context, req sns.SnsSendImageRequest) (*sns.SnsSendImageResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&sns.SnsSendImageResponse{}).Post(snsSendImage)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用snsSendImage方法失败")
		return nil, err
	}

	return resp.Result().(*sns.SnsSendImageResponse), nil
}
