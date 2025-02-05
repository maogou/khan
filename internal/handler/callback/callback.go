package callback

import (
	"smallBot/internal/handler/callback/chain"
	"smallBot/internal/sdk/khan"
)

type CallbackHandler struct {
	sdk   *khan.Khan
	chain chain.Chain
}

func NewCallbackHandler(sdk *khan.Khan, chain chain.Chain) *CallbackHandler {
	return &CallbackHandler{
		sdk:   sdk,
		chain: chain,
	}
}
