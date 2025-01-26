package transform

type ContactGetMobileRequest struct {
	Appid   string   `json:"appid"`
	PhoneNo []string `json:"phone_no"`
}
