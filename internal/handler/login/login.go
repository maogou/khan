package login

import "smallBot/internal/sdk/khan"

type LoginHandler struct {
	sdk *khan.Khan
}

func NewLoginHandler(sdk *khan.Khan) *LoginHandler {
	return &LoginHandler{
		sdk: sdk,
	}
}
