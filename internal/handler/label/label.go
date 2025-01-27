package label

import "smallBot/internal/sdk/gewe"

type LabelHandler struct {
	sdk *gewe.Gewe
}

func NewLabelHandler(sdk *gewe.Gewe) *LabelHandler {
	return &LabelHandler{
		sdk: sdk,
	}
}
