package chain

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/sdk/gewe"

	"github.com/rs/zerolog/log"
)

type Reconnection struct {
	BaseHandler
	sdk *gewe.Gewe
}

func (r *Reconnection) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.Info().Msg("调用Reconnection->HandlerRequest方法")

}

func (r *Reconnection) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	//TODO implement me
	return true
}

func (r *Reconnection) Process(ctx context.Context, param v1.CollectRequest) {
	//TODO implement me
	log.Info().Msg("调用Reconnection->Process方法")
}

func NewReconnection(sdk *gewe.Gewe) *Reconnection {
	return &Reconnection{
		sdk: sdk,
	}
}

var _ Chain = (*Reconnection)(nil)
