package message

import (
	"maogou/khan/internal/sdk/khan"
)

type MessageHandler struct {
	sdk *khan.Khan
}

func NewMessageHandler(sdk *khan.Khan) *MessageHandler {
	return &MessageHandler{
		sdk: sdk,
	}
}
