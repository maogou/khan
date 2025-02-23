package processing

import (
	"context"
	"smallBot/internal/sdk/khan"
	"smallBot/internal/task/forward/chain"
)

type Processing struct {
	handlers []chain.Handler
}

func (pt *Processing) Execute(ctx context.Context, appId string) error {
	var pld = chain.PipelineData{
		AppID: appId,
	}

	for _, handler := range pt.handlers {
		if err := handler.Handle(ctx, &pld); err != nil {
			return err
		}

	}
	return nil
}

func NewForwardChain(sdk *khan.Khan) *Processing {
	return &Processing{
		handlers: []chain.Handler{
			chain.NewVerifyHandler(sdk),
			chain.NewSyncHandler(sdk),
			chain.NewFavIdHandler(sdk),
			chain.NewFavDetailHandler(sdk),
			chain.NewParseFavHandler(sdk),
			chain.NewSnsDetailHandler(sdk),
			chain.NewPareSnsHandler(sdk),
			chain.NewSendHandler(sdk),
			chain.NewDeleteFavHandler(sdk),
			chain.NewNotifyHandler(sdk),
		},
	}
}
