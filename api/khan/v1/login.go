package v1

type GetLoginQrCodeRequest struct {
	AppId    string `json:"appId" binding:"required"`
	ProxyIp  string `json:"proxyIp"`
	RegionId string `json:"regionId"`
}

type GetLoginQrCodeResponse struct {
	AppId       string `json:"appId"`
	QrData      string `json:"qrData"`
	QrImgBase64 string `json:"qrImgBase64"`
	Uuid        string `json:"uuid"`
	NewKey      string `json:"newKey"`
}

type CheckLoginRequest struct {
	AppId  string `json:"appId" binding:"required"`
	Uuid   string `json:"uuid" binding:"required"`
	Newkey string `json:"newKey" binding:"required"`
}

type CheckLoginLoginInfo struct {
	Uin      int    `json:"uin"`
	Wxid     string `json:"wxid"`
	NickName string `json:"nickName"`
	Mobile   string `json:"mobile"`
	Alias    string `json:"alias"`
}

type CheckLoginResponse struct {
	Uuid        string              `json:"uuid"`
	HeadImgUrl  string              `json:"headImgUrl"`
	NickName    string              `json:"nickName"`
	ExpiredTime int                 `json:"expiredTime"`
	Status      int                 `json:"status"`
	LoginInfo   CheckLoginLoginInfo `json:"loginInfo"`
}

type LoginCreateApp struct {
	AppId string `json:"app_id"`
	Tip   string `json:"tip"`
}

type LogoutRequest struct {
	AppId string `json:"appId" binding:"required"`
}

type CheckOnlineRequest struct {
	AppId string `json:"appId" binding:"required"`
}

type AppIdRequest struct {
	AppId string `json:"appId" binding:"required"`
}
