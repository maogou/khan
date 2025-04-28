package setting

import "maogou/khan/internal/sdk/khan"

type SettingHandler struct {
	sdk *khan.Khan
}

func NewSettingHandler(sdk *khan.Khan) *SettingHandler {
	return &SettingHandler{
		sdk: sdk,
	}
}
