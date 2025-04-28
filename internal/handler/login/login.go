package login

import "maogou/khan/internal/sdk/khan"

type LoginHandler struct {
	sdk *khan.Khan
}

func NewLoginHandler(sdk *khan.Khan) *LoginHandler {
	return &LoginHandler{
		sdk: sdk,
	}
}
