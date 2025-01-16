package message

import (
	"smallBot/internal/sdk/gewe"
)

type MessageHandler struct {
	sdk *gewe.Gewe
}

func NewMessageHandler(sdk *gewe.Gewe) *MessageHandler {
	return &MessageHandler{
		sdk: sdk,
	}
}
