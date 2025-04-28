package favor

import "maogou/khan/internal/sdk/khan"

type FavorHandler struct {
	sdk *khan.Khan
}

func NewFavorHandler(sdk *khan.Khan) *FavorHandler {
	return &FavorHandler{
		sdk: sdk,
	}
}
