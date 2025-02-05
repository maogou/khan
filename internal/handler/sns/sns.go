package sns

import "smallBot/internal/sdk/khan"

type SnsHandler struct {
	sdk *khan.Khan
}

func NewSnsHandler(skd *khan.Khan) *SnsHandler {
	return &SnsHandler{
		sdk: skd,
	}
}
