package chain

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
)

type Verify struct {
	BaseHandler
	sdk *khan.Khan
}

func (v *Verify) HandlerRequest(ctx context.Context, param v1.CollectRequest) {

	log.C(ctx).Info().Msg("调用Verify->HandlerRequest方法")

	if v.IsCanHandler(ctx, param) {
		v.ExecuteNext(ctx, param)
	}
}

func (v *Verify) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用Verify->IsCanHandler方法")

	_, err := help.Permission(ctx, v.sdk.Rdb())

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("License已过期")
		return false
	}

	return true
}

func (v *Verify) Process(ctx context.Context, param v1.CollectRequest) error {

	return nil
}

var _ Chain = (*Verify)(nil)

func NewVerify(sdk *khan.Khan) *Verify {
	return &Verify{
		sdk: sdk,
	}
}
