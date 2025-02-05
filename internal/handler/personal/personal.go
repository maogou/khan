package personal

import "smallBot/internal/sdk/khan"

type PersonalHandler struct {
	sdk *khan.Khan
}

func NewPersonalHandler(sdk *khan.Khan) *PersonalHandler {
	return &PersonalHandler{
		sdk: sdk,
	}
}
