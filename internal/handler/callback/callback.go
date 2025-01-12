package callback

import (
	"smallBot/internal/handler/callback/chain"
	"smallBot/internal/sdk/gewe"
)

type CallbackHandler struct {
	sdk   *gewe.Gewe
	chain chain.Chain
}

func NewCallbackHandler(sdk *gewe.Gewe, chain chain.Chain) *CallbackHandler {
	return &CallbackHandler{
		sdk:   sdk,
		chain: chain,
	}
}
