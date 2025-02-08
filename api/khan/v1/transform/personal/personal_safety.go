package personal

type PersonalSafetyRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalSafetyResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Info struct {
			Devicelist []struct {
				Uuid       string `json:"uuid"`
				Devicename string `json:"devicename"`
				Devicetype string `json:"devicetype"`
				Lasttime   int    `json:"lasttime"`
			} `json:"devicelist"`
			BHasVoice    bool `json:"bHasVoice"`
			BSwitchVoice bool `json:"bSwitchVoice"`
			BHasFace     bool `json:"bHasFace"`
			BSwitchFace  bool `json:"bSwitchFace"`
			BHasWxPwd    bool `json:"bHasWxPwd"`
		} `json:"info"`
	} `json:"data"`
}
