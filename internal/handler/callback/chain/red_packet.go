package chain

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/sdk/khan"
)

type RedPacket struct {
	BaseHandler
	sdk *khan.Khan
}

func (r RedPacket) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	//TODO implement me
	panic("implement me")
}

func (r RedPacket) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	//TODO implement me
	panic("implement me")
}

func (r RedPacket) Process(ctx context.Context, param v1.CollectRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewRedPacket(sdk *khan.Khan) *RedPacket {
	return &RedPacket{
		sdk: sdk,
	}
}

var _ Chain = (*RedPacket)(nil)
