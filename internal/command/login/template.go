package login

type template struct {
	b BootStrap
}

func (t *template) Start(appId string) {

	if len(appId) == 0 {
		t.b.SetAppId(appId)
	}

	if t.b.CanNext() {
		t.b.GetLoginQrCode()
	}

	if t.b.CanNext() {
		t.b.PrintQrCode()
	}

	if t.b.CanNext() {
		t.b.Confirm()
	}

	if t.b.CanNext() {
		t.b.CheckLogin()
	}

	if t.b.CanNext() {
		t.b.PrintConfig()
	}

	if t.b.CanNext() {
		t.b.Welcome()
	}
}
