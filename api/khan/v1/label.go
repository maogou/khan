package v1

type LabelAddRequest struct {
	AppId     string `json:"appId" binding:"required"`
	LabelName string `json:"labelName" binding:"required"`
}

type LabelAddResponse struct {
	LabelName string `json:"labelName"`
	LabelId   int    `json:"labelId"`
}

type LabelDeleteRequest struct {
	AppId    string `json:"appId" binding:"required"`
	LabelIds string `json:"labelIds" binding:"required"`
}

type LabelListRequest struct {
	AppId string `json:"appId" binding:"required"`
}

type LabelListItem struct {
	LabelName string `json:"labelName"`
	LabelId   int    `json:"labelId"`
}

type LabelListResponse struct {
	LabelList []LabelListItem `json:"labelList"`
}

type LabelModifyRequest struct {
	AppId    string   `json:"appId" binding:"required"`
	LabelIds string   `json:"labelIds" binding:"required"`
	WxIds    []string `json:"wxIds" binding:"required"`
}
