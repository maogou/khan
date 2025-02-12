package v1

type Permission struct {
	Permission map[string]int `json:"permission"`
	Wid        []string       `json:"wid"`
}
