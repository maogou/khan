package login

import "smallBot/internal/sdk/gewe"

type LoginHandler struct {
	sdk *gewe.Gewe
}

func NewLoginHandler(sdk *gewe.Gewe) *LoginHandler {
	return &LoginHandler{
		sdk: sdk,
	}
}
