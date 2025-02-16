package v1

type Permission struct {
	Permission map[string]int    `json:"permission"`
	AppId      []string          `json:"appid"`
	Token      map[string]string `json:"token"`
}
