package callback

import (
	"maogou/khan/internal/handler/callback/chain"
	"maogou/khan/internal/sdk/khan"
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
