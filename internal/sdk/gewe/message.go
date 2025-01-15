package gewe

import (
	"context"
	v1 "smallBot/api/gewe/v1"
)

func (g *Gewe) PostText(ctx context.Context, req v1.PostTextRequest) (*v1.PostTextResponse, error) {
	resp, err := g.client.R().SetBody(req).SetBody(&v1.PostTextResponse{}).Post(postText)

	if err != nil {
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
