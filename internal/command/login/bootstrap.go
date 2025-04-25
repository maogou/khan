package login

type BootStrap interface {
	SetAppId(appId string)
	GetLoginQrCode()
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
