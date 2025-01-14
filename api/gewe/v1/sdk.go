package v1

type CommonResponse struct {
	Ret    int    `json:"ret"`
	Msg    string `json:"msg"`
	MsgErr string `json:"msg_err,omitempty"`
}

type LoginQrCodeRequest struct {
	AppId string `json:"appId" validate:"required"`
}

type LoginQrCodeData struct {
	Base64 string `json:"base64"`
	Url    string `json:"url"`
	Uuid   string `json:"uuid"`
	NKey   string `json:"nkey"`
	Id     string `json:"id"`
}

type LoginQrCodeResponse struct {
	CommonResponse
	Data LoginQrCodeData `json:"data"`
}

type CheckLoginQrCodeRequest struct {
	AppId string `json:"appid" validate:"required"`
	NKey  string `json:"nkey" validate:"required"`
	Uuid  string `json:"uuid" validate:"required"`
}

type CheckLoginQrCodeStatusInfo struct {
	Uuid                    string `json:"uuid"`
	Status                  int    `json:"status"`
	HeadImgUrl              string `json:"headImgUrl"`
	PushLoginUrlexpiredTime int    `json:"pushLoginUrlexpiredTime"`
	NickName                string `json:"nickName"`
	ExpiredTime             int    `json:"expiredTime"`
}

type CheckLoginQrCodeData struct {
	StatusInfo CheckLoginQrCodeStatusInfo `json:"status_info"`
}

type CheckLoginQrCodeResponse struct {
	CommonResponse
	Data CheckLoginQrCodeData `json:"data"`
}

type LongOpenRequest struct {
	AppId      string `json:"appid" validate:"required"`
	CleanCache bool   `json:"cleanCache" validate:"required"`
	Host       string `json:"host" validate:"required"`
	Timeout    int    `json:"timeout" validate:"required"`
}

type LongOpenResponse struct {
	CommonResponse
	Data bool `json:"data"`
}

type SecAutoAuthRequest struct {
	AppId  string `json:"appid" validate:"required"`
	SdkVer string `json:"sdk_ver" validate:"required"`
}

type SecAutoAuthBaseResponse struct {
	Ret    int `json:"ret"`
	ErrMsg struct {
		String string `json:"string"`
	} `json:"errMsg"`
}

type SecAutoAuthResponseData struct {
	BaseResponse SecAutoAuthBaseResponse `json:"baseResponse"`
}

type SecAutoAuthResponse struct {
	CommonResponse
	Data SecAutoAuthResponseData `json:"data"`
}

type CreateAppRequest struct {
	Model      string `json:"model"`
	SdkVer     string `json:"sdk_ver"`
	DeviceName string `json:"devicename"`
	Country    string `json:"country"`
}

type CreateAppData struct {
	AppId string `json:"appid"`
	Id    string `json:"id"`
	Desc  string `json:"desc"`
}

type CreateAppResponse struct {
	CommonResponse
	Data CreateAppData `json:"data"`
}

type HearBeatRequest struct {
	AppId string `json:"appid" validate:"required"`
}

type HearBeatResponse struct {
	CommonResponse
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		NextTime                  int   `json:"NextTime"`
		Selector                  int64 `json:"Selector"`
		BlueToothBroadCastContent struct {
			ILen int `json:"iLen"`
		} `json:"BlueToothBroadCastContent"`
	} `json:"data"`
}
