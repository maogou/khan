package gewe

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/log"
)

func (g *Gewe) PostText(ctx context.Context, req v1.PostTextRequest) (*v1.PostTextResponse, error) {
	resp, err := g.client.R().SetBody(req).SetResult(&v1.PostTextResponse{}).Post(postText)
	if err != nil {
		log.C(ctx).Error().Msg("调用PostText方法失败")
		return nil, err
	}

	return resp.Result().(*v1.PostTextResponse), nil
}

func (g *Gewe) PushMsg(ctx context.Context, req v1.CollectRequest, callback string) error {
	_, err := g.client.R().SetBody(req).Post(callback)

	if err != nil {
		return err
	}

	return nil
}
