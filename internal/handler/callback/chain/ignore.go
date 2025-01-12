package chain

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/sdk/gewe"

	"github.com/rs/zerolog/log"
)

type Ignore struct {
	BaseHandler
	sdk *gewe.Gewe
}

func NewIgnore(sdk *gewe.Gewe) *Ignore {
	return &Ignore{
		sdk: sdk,
	}
}

var _ Chain = (*Ignore)(nil)

func (i *Ignore) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.Info().Msg("调用Ignore->HandlerRequest方法")

	if i.IsCanHandler(ctx, param) {
		i.Process(ctx, param)
	}

	i.NextHandler.HandlerRequest(ctx, param)
}

func (i *Ignore) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	if len(param.AppId) == 0 {
		return false
	}

	return true
}

func (i *Ignore) Process(ctx context.Context, param v1.CollectRequest) {
	//TODO implement me
	log.Info().Msg("调用Ignore->Process方法")
}
