package chain

import (
	"context"
	v1 "smallBot/api/gewe/v1"
)

type Chain interface {
	SetNext(handler Chain)
	HandlerRequest(ctx context.Context, param v1.CollectRequest)
	IsCanHandler(ctx context.Context, param v1.CollectRequest) bool
	Process(ctx context.Context, param v1.CollectRequest) error
}

type BaseHandler struct {
	NextHandler Chain
}

func (h *BaseHandler) SetNext(handler Chain) {
	h.NextHandler = handler
}
