package personal

import "smallBot/internal/sdk/gewe"

type PersonalHandler struct {
	sdk *gewe.Gewe
}

func NewPersonalHandler(sdk *gewe.Gewe) *PersonalHandler {
	return &PersonalHandler{
		sdk: sdk,
	}
}
