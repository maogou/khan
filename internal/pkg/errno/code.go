package errno

var (
	Ok                  = &ErrNo{Code: 0, Message: "OK"}
	PageNotFoundError   = &ErrNo{Code: 404, Message: "访问的页面不存在"}
	MethodNotAllowError = &ErrNo{Code: 405, Message: "访问的路由请求方式错误"}
	InternalServerError = &ErrNo{Code: 10001, Message: "服务器内部,请稍后再试"}
	ValidateError       = &ErrNo{Code: 10002, Message: "参数验证错误"}

	SendMsgError                = &ErrNo{Code: 10003, Message: "发送消息失败"}
	DownloadImgError            = &ErrNo{Code: 10004, Message: "下载图片失败"}
	DownloadNameEmptyError      = &ErrNo{Code: 10005, Message: "下载文件名不能为空"}
	DownloadFileNotExistError   = &ErrNo{Code: 10006, Message: "下载文件不存在"}
	DownloadFileOpenError       = &ErrNo{Code: 10007, Message: "文件打开失败"}
	DownloadFileStatError       = &ErrNo{Code: 10008, Message: "获取文件信息失败"}
	PersonalProfileError        = &ErrNo{Code: 10009, Message: "获取个人资料失败"}
	PersonalQrcodeError         = &ErrNo{Code: 10010, Message: "获取个人二维码失败"}
	PersonalSafetyError         = &ErrNo{Code: 10011, Message: "获取个人安全设备失败"}
	PersonalPrivacySettingError = &ErrNo{Code: 10012, Message: "设置个人隐私设置失败"}
	PersonalUploadHeadImgError  = &ErrNo{Code: 10013, Message: "上传头像失败"}
	PostTextError               = &ErrNo{Code: 10014, Message: "发送文本消息失败"}
	PostImageError              = &ErrNo{Code: 10015, Message: "发送图片消息失败"}
	PostFileError               = &ErrNo{Code: 10016, Message: "发送文件消息失败"}
	PostLinkError               = &ErrNo{Code: 10017, Message: "发送链接消息失败"}
	PostVoiceError              = &ErrNo{Code: 10018, Message: "发送语音消息失败"}
	PostVideoError              = &ErrNo{Code: 10019, Message: "发送视频消息失败"}
	PostNameCardError           = &ErrNo{Code: 10020, Message: "发送名片消息失败"}
	XmlEncodeError              = &ErrNo{Code: 10021, Message: "xml编码失败"}
	PostEmojiError              = &ErrNo{Code: 10022, Message: "发送表情消息失败"}

	ContactListError        = &ErrNo{Code: 10023, Message: "获取联系人列表失败"}
	RequestCanceled         = &ErrNo{Code: 10024, Message: "请求被取消"}
	ContactSearchError      = &ErrNo{Code: 10025, Message: "搜索联系人失败"}
	ContactAddError         = &ErrNo{Code: 10026, Message: "添加联系人失败"}
	ContactDelError         = &ErrNo{Code: 10027, Message: "删除联系人失败"}
	ContactInfoError        = &ErrNo{Code: 10028, Message: "获取联系人信息失败"}
	ContactDetailError      = &ErrNo{Code: 10029, Message: "获取联系人详细信息失败"}
	ContactSetOnlyChatError = &ErrNo{Code: 10030, Message: "设置联系人为仅聊天失败"}
	ContactSetRemarkError   = &ErrNo{Code: 10031, Message: "设置联系人备注失败"}
)
