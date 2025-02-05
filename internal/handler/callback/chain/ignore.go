package chain

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
)

type Ignore struct {
	BaseHandler
	sdk *khan.Khan
}

func NewIgnore(sdk *khan.Khan) *Ignore {
	return &Ignore{
		sdk: sdk,
	}
}

var _ Chain = (*Ignore)(nil)

func (i *Ignore) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.C(ctx).Info().Msg("调用Ignore->HandlerRequest方法")

	if i.IsCanHandler(ctx, param) {
		if i.NextHandler != nil {
			i.NextHandler.HandlerRequest(ctx, param)
		}
	}
}

func (i *Ignore) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用Ignore->IsCanHandler方法")
	if len(param.AppId) == 0 {
		return false
	}

	msgType := param.TypeName

	if len(msgType) == 0 {
		return false
	}

	return true

}

func (i *Ignore) Process(ctx context.Context, param v1.CollectRequest) error {
	log.C(ctx).Info().Msg("调用Ignore->Process方法")

	return nil
}
