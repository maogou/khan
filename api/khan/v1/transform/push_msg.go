package transform

type CallbackRequest struct {
	AppId    string `json:"Appid" validate:"required"`
	TypeName string `json:"TypeName" validate:"required"`
	Data     any    `json:"Data" validate:"required"`
}
