package transform

type ContactUploadMobileRequest struct {
	Appid   string   `json:"appid"`
	Opcode  int      `json:"opcode"`
	PhoneNo []string `json:"phone_no"`
}
