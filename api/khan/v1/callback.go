package v1

import "encoding/json"

type CollectRequest struct {
	AppId    string          `json:"appid" validate:"required"`
	TypeName string          `json:"type_name" validate:"required"`
	Data     json.RawMessage `json:"data" validate:"required"`
}

type SetCallbackRequest struct {
	AppId       string `json:"appid" validate:"required"`
	CallbackUrl string `json:"callbackUrl" validate:"required"`
}
