package label

import "smallBot/internal/sdk/khan"

type LabelHandler struct {
	sdk *khan.Khan
}

func NewLabelHandler(sdk *khan.Khan) *LabelHandler {
	return &LabelHandler{
		sdk: sdk,
	}
}
