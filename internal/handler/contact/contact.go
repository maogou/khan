package contact

import "smallBot/internal/sdk/gewe"

type ContactHandler struct {
	sdk *gewe.Gewe
}

func NewContactHandler(sdk *gewe.Gewe) *ContactHandler {
	return &ContactHandler{
		sdk: sdk,
	}
}
