package v1

type CollectRequest struct {
	AppId    string `json:"app_id" validate:"required"`
	TypeName string `json:"type_name" validate:"required"`
	Data     string `json:"data" validate:"required"`
}
