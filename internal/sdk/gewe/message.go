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

func (g *Gewe) PushMsg(ctx context.Context, req v1.CollectRequest, callback string) error {
	_, err := g.client.R().SetBody(req).Post(callback)

	if err != nil {
		return err
	}

	return nil
}
