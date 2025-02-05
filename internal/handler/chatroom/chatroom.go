package chatroom

import "smallBot/internal/sdk/khan"

type ChatRoomHandler struct {
	sdk *khan.Khan
}

func NewChatRoomHandler(sdk *khan.Khan) *ChatRoomHandler {
	return &ChatRoomHandler{
		sdk: sdk,
	}
}
