package message

import "encoding/json"

type CallbackRequest struct {
	AppId    string          `json:"Appid" validate:"required"`
	TypeName string          `json:"TypeName" validate:"required"`
	Data     json.RawMessage `json:"Data" validate:"required"`
	Wxid     string          `json:"Wxid" validate:"required"`
}
