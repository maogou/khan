package chain

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/log"
)

type Verify struct {
	BaseHandler
	auth *license.License
}

func (v *Verify) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.C(ctx).Info().Any("param", param).Msg("调用Verify->HandlerRequest方法")

	if v.IsCanHandler(ctx, param) {
		if v.NextHandler != nil {
			v.NextHandler.HandlerRequest(ctx, param)
		}
	}
}

func (v *Verify) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用Verify->IsCanHandler方法")
	if v.auth.Expired() {
		log.C(ctx).Warn().Msg("License已过期")
		return false
	}

	return true
}

func (v *Verify) Process(ctx context.Context, param v1.CollectRequest) error {

	return nil
}

var _ Chain = (*Verify)(nil)

func NewVerify(l *license.License) *Verify {
	return &Verify{
		auth: l,
	}
}
