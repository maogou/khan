package login

type BootStrap interface {
	GetAppId()
	GetLoginQrCode(appId string)
	PrintQrCode()
	Confirm()
	CheckLogin()
	PrintConfig()
	Welcome()
	CanNext() bool
}
