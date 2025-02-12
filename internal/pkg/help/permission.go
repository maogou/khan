package help

import "encoding/json"

func AccessVipPermission() map[string]int {
	return map[string]int{
		"group":    1,
		"contact":  1,
		"favor":    1,
		"label":    1,
		"message":  1,
		"personal": 1,
		"download": 1,
		"sns":      1,
		"transfer": 1,
	}
}

func AccessPermission() map[string]int {
	return map[string]int{
		"group":    1,
		"contact":  1,
		"favor":    1,
		"label":    1,
		"message":  1,
		"personal": 1,
		"download": 1,
		"sns":      0,
		"transfer": 0,
	}
}

func TransferPermissionToByte(permission map[string]int) ([]byte, error) {
	return json.Marshal(permission)
}
