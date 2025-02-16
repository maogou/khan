package chain

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/sdk/khan"
)

type CollectMoney struct {
	BaseHandler
	sdk *khan.Khan
}

func (c CollectMoney) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	//TODO implement me
	panic("implement me")
}

func (c CollectMoney) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	//TODO implement me
	panic("implement me")
}

func (c CollectMoney) Process(ctx context.Context, param v1.CollectRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewCollectMoney(sdk *khan.Khan) *CollectMoney {
	return &CollectMoney{
		sdk: sdk,
	}
}

var _ Chain = (*CollectMoney)(nil)
