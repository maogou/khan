package login

type BootStrap interface {
	GetAppId()
	GetLoginQrCode(appId string)
	PrintQrCode()
	Confirm()
	CheckLogin()
	PrintConfig()
	HearBeat()
	AutoAuth()
	Open()
	Welcome()
	CanNext() bool
}
