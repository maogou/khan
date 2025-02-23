package chain

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/favor"
	"smallBot/api/khan/v1/transform/sns"
)

type Handler interface {
	Handle(ctx context.Context, pld *PipelineData) error
	SetNext(handler Handler)
}

type PipelineData struct {
	AppID        string
	FavID        int
	SnsID        string
	ContentStyle int
	IsForwardOk  bool
	FavDetail    *favor.FavorDetailResponse
	SnsDetail    *sns.SnsDetailResponse
	Req          v1.SnsForwardRequest
	Timeline     v1.TimelineObject
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}
