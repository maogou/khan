package favor

import "smallBot/internal/sdk/khan"

type FavorHandler struct {
	sdk *khan.Khan
}

func NewFavorHandler(sdk *khan.Khan) *FavorHandler {
	return &FavorHandler{
		sdk: sdk,
	}
}
