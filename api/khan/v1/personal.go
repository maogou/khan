package v1

type PersonalProfileRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalProfileResponse struct {
	Alias           string `json:"alias"`
	Wxid            string `json:"wxid"`
	NickName        string `json:"nickName"`
	Mobile          string `json:"mobile"`
	Uin             int    `json:"uin"`
	Sex             int    `json:"sex"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Signature       string `json:"signature"`
	Country         string `json:"country"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	RegCountry      string `json:"regCountry"`
	SnsBgImg        string `json:"snsBgImg"`
}

type PersonalQrcodeRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalQrcodeResponse struct {
	QrCode string `json:"qrCode"`
}

type PersonalSafetyRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalSafetyItem struct {
	Uuid       string `json:"uuid"`
	DeviceName string `json:"deviceName"`
	DeviceType string `json:"deviceType"`
	LastTime   int    `json:"lastTime"`
}

type PersonalSafetResponse struct {
	List []PersonalSafetyItem `json:"list"`
}

type PersonalPrivacySettingRequest struct {
	AppId  string `json:"appId" binding:"required"`
	Open   bool   `json:"open" binding:"required"`
	Option int    `json:"option" binding:"required"`
}

type PersonalUploadHeadImgRequest struct {
	AppId      string `json:"appId" binding:"required"`
	HeadImgUrl string `json:"headImgUrl" binding:"required"`
}

type PersonalLicenseResponse struct {
	AppId          []string                    `json:"app_id"`
	Token          map[string]string           `json:"token"`
	CreateTime     string                      `json:"create_time"`
	ExpireTime     string                      `json:"expire_time"`
	RemainDays     int                         `json:"remain_days"`
	LicenseKey     string                      `json:"license_key"`
	LicenseContent string                      `json:"license_content"`
	Permission     []PersonalLicensePermission `json:"permission"`
	LicenseType    string                      `json:"license_type"`
}

type PersonalLicensePermission struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
