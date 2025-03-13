package v1

type LicenseUploadRequest struct {
	File string `form:"file" binding:"required"`
}

type AppIdsRequest struct {
	AppIdSwitch []AppIdSwitchItem `json:"appIdSwitch" binding:"required,dive"`
	Func        int               `json:"func" binding:"required,oneof=1 2 3"`
}

type AppIdSwitchItem struct {
	AppId  string `json:"appId" binding:"required"`
	Switch *int   `json:"switch" binding:"required,oneof=0 1"`
}

type SuccessFailResponse struct {
	Success []string `json:"success"`
	Fail    []string `json:"fail"`
}
