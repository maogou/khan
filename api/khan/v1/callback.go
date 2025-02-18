package v1

import "encoding/json"

type CollectRequest struct {
	AppId    string          `json:"appid" validate:"required"`
	TypeName string          `json:"type_name" validate:"required"`
	Data     json.RawMessage `json:"data" validate:"required"`
	MsgType  int             `json:"MsgType"`
}

type SetCallbackRequest struct {
	Token       string `json:"token" validate:"required"`
	CallbackUrl string `json:"callbackUrl" validate:"required"`
}

type CallbackMessageType struct {
	MsgType int `json:"MsgType"`
}
