package v1

type ResponseCommon struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

type ResponseSuccessWithoutData struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}
