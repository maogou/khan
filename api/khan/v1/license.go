package v1

type LicensePermission struct {
	Permission map[string]map[string]int `json:"permission"`
}

type LicenseYourselfWidRequest struct {
	YourselfWXid string `json:"yourself_wxid" binding:"required"`
}
