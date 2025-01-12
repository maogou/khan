package v1

type CollectRequest struct {
	AppId    string `json:"appid" validate:"required"`
	TypeName string `json:"type_name" validate:"required"`
	Data     any    `json:"data" validate:"required"`
}
