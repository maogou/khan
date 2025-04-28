package contact

import "maogou/khan/internal/sdk/khan"

type ContactHandler struct {
	sdk *khan.Khan
}

func NewContactHandler(sdk *khan.Khan) *ContactHandler {
	return &ContactHandler{
		sdk: sdk,
	}
}
