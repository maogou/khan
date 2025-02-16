package help

func AccessVipPermission() map[string]int {
	return map[string]int{
		"login":        1,
		"group":        1,
		"contact":      1,
		"favor":        1,
		"label":        1,
		"message":      1,
		"personal":     1,
		"download":     1,
		"sns":          1,
		"transfer":     1,
		"finder":       0,
		"redPacket":    0,
		"collectMoney": 0,
	}
}

func AccessPermission() map[string]int {
	return map[string]int{
		"login":        1,
		"group":        1,
		"contact":      1,
		"favor":        1,
		"label":        1,
		"message":      1,
		"personal":     1,
		"download":     1,
		"sns":          0,
		"transfer":     0,
		"finder":       0,
		"redPacket":    0,
		"collectMoney": 0,
	}
}

func AccessPermissionName() map[string]string {
	return map[string]string{
		"group":        "群模块",
		"contact":      "联系人模块",
		"favor":        "收藏夹模块",
		"label":        "标签模块",
		"message":      "消息模块",
		"personal":     "个人模块",
		"download":     "下载模块",
		"sns":          "朋友圈模块",
		"transfer":     "迁移模块",
		"login":        "登录模块",
		"finder":       "视频号模块",
		"redPacket":    "自动🧧",
		"collectMoney": "收款💰",
	}

}

func TransferPermissionToName(permission map[string]int) []string {
	names := make([]string, 0)
	permissionName := AccessPermissionName()
	for k, v := range permission {
		if v == 0 {
			continue
		}

		if name, ok := permissionName[k]; ok {
			names = append(names, name)
		}
	}

	return names
}
