package callback

import "smallBot/internal/sdk/gewe"

type CallbackHandler struct {
	sdk *gewe.Gewe
}

func NewCallbackHandler(sdk *gewe.Gewe) *CallbackHandler {
	return &CallbackHandler{
		sdk: sdk,
	}
}
