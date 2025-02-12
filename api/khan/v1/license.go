package v1

type LicensePermission struct {
	Permission map[string]map[string]int `json:"permission"`
}
