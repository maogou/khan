package errno

var (
	Ok                  = &ErrNo{Code: 0, Message: "OK"}
	PageNotFoundError   = &ErrNo{Code: 404, Message: "访问的页面不存在"}
	MethodNotAllowError = &ErrNo{Code: 405, Message: "访问的路由请求方式错误"}
	RedisGetDataError   = &ErrNo{Code: 406, Message: "redis获取数据失败"}
	RedisNoDataError    = &ErrNo{Code: 407, Message: "redis获取数据为空"}
	TemplateParseError  = &ErrNo{Code: 408, Message: "模板解析错误"}
	ConvertError        = &ErrNo{Code: 409, Message: "类型转换错误"}
	DecryptTokenFail    = &ErrNo{Code: 410, Message: "token解析失败"}
	CreateAppErr        = &ErrNo{Code: 411, Message: "创建应用失败"}
	CallbackMsgTypeErr  = &ErrNo{Code: 412, Message: "回调消息类型解析错误"}
	GetFileBase64Error  = &ErrNo{Code: 10000, Message: "获取文件base64失败"}
	InternalServerError = &ErrNo{Code: 10001, Message: "服务器内部,请稍后再试"}
	ValidateError       = &ErrNo{Code: 10002, Message: "参数验证错误"}
	ForwardMiniAppError = &ErrNo{Code: 10002, Message: "转发小程序失败"}

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
	RevokeMsgError              = &ErrNo{Code: 10021, Message: "撤回消息失败"}
	ForwardUrlError             = &ErrNo{Code: 10022, Message: "转发url失败"}
	ForwardVideoError           = &ErrNo{Code: 10022, Message: "转发视频失败"}
	ForwardFileError            = &ErrNo{Code: 10022, Message: "转发文件失败"}
	SendMiniAppError            = &ErrNo{Code: 10022, Message: "发送小程序失败"}

	ContactListError        = &ErrNo{Code: 10023, Message: "获取联系人列表失败"}
	RequestCanceled         = &ErrNo{Code: 10024, Message: "请求被取消"}
	ContactSearchError      = &ErrNo{Code: 10025, Message: "搜索联系人失败"}
	ContactAddError         = &ErrNo{Code: 10026, Message: "添加联系人失败"}
	ContactDelError         = &ErrNo{Code: 10027, Message: "删除联系人失败"}
	ContactInfoError        = &ErrNo{Code: 10028, Message: "获取联系人信息失败"}
	ContactDetailError      = &ErrNo{Code: 10029, Message: "获取联系人详细信息失败"}
	ContactSetOnlyChatError = &ErrNo{Code: 10030, Message: "设置联系人为仅聊天失败"}
	ContactSetRemarkError   = &ErrNo{Code: 10031, Message: "设置联系人备注失败"}

	LabelAddError    = &ErrNo{Code: 10032, Message: "添加标签失败"}
	LabelDeleteError = &ErrNo{Code: 10033, Message: "删除标签失败"}
	LabelListError   = &ErrNo{Code: 10034, Message: "获取标签列表失败"}
	LabelModifyError = &ErrNo{Code: 10035, Message: "修改标签失败"}

	LoginQrCodeError              = &ErrNo{Code: 10036, Message: "获取登录二维码失败"}
	CheckLoginQrCodeError         = &ErrNo{Code: 10037, Message: "检查登录二维码失败"}
	SetCallbackError              = &ErrNo{Code: 10038, Message: "请检查回调地址,此回调地址容器内无法访问"}
	SetTokenCallbackError         = &ErrNo{Code: 10039, Message: "设置token对应的回调地址失败"}
	SetWxLoginCacheError          = &ErrNo{Code: 10040, Message: "设置wx登录缓存失败"}
	SetWxLoginQrCodeCacheError    = &ErrNo{Code: 10041, Message: "设置wx登录二维码缓存失败"}
	GetWxLoginQrCodeCacheError    = &ErrNo{Code: 10042, Message: "获取wx登录二维码缓存失败"}
	ExpireWxLoginQrCodeCacheError = &ErrNo{Code: 10043, Message: "登录二维码缓存已过期,请重新获取登录二维码"}
	ChatroomCreateError           = &ErrNo{Code: 10039, Message: "创建群聊失败"}
	ChatroomModifyNameError       = &ErrNo{Code: 10040, Message: "修改群聊名称失败"}
	ChatroomModifyRemarkError     = &ErrNo{Code: 10041, Message: "修改群聊备注失败"}
	ChatroomModifyMyselfNameError = &ErrNo{Code: 10042, Message: "修改群聊名称失败"}
	ChatroomSetAnnouncementError  = &ErrNo{Code: 10043, Message: "修改群聊公告失败"}
	ChatroomGetAnnouncementError  = &ErrNo{Code: 10044, Message: "获取群聊公告失败"}
	ChatroomDetailError           = &ErrNo{Code: 10045, Message: "获取群聊详情失败"}
	XmlDecodeError                = &ErrNo{Code: 10046, Message: "xml解码失败"}
	ExecPathError                 = &ErrNo{Code: 10047, Message: "获取执行路径失败"}

	FavorSyncError   = &ErrNo{Code: 10048, Message: "同步收藏失败"}
	FavorDetailError = &ErrNo{Code: 10049, Message: "获取收藏详情失败"}
	FavorDeleteError = &ErrNo{Code: 10050, Message: "删除收藏失败"}

	ChatroomMemberError     = &ErrNo{Code: 10051, Message: "获取群聊成员失败"}
	ChatroomTopPingError    = &ErrNo{Code: 10052, Message: "置顶群聊失败"}
	ChatroomSetSilenceError = &ErrNo{Code: 10053, Message: "设置群聊静音失败"}
	ChatroomQrcodeError     = &ErrNo{Code: 10054, Message: "获取群聊二维码失败"}

	SnsDetailError                    = &ErrNo{Code: 10055, Message: "获取朋友圈详情失败"}
	SnsMyselfPageError                = &ErrNo{Code: 10056, Message: "获取自己的朋友圈列表失败"}
	SnsFriendPageError                = &ErrNo{Code: 10057, Message: "获取好友朋友圈列表失败"}
	SnsLikeError                      = &ErrNo{Code: 10058, Message: "点赞朋友圈失败"}
	SnsCancelLikeError                = &ErrNo{Code: 10059, Message: "取消点赞朋友圈失败"}
	SnsCommentError                   = &ErrNo{Code: 10060, Message: "评论朋友圈失败"}
	SnsDeleteCommentError             = &ErrNo{Code: 10061, Message: "删除朋友圈评论失败"}
	SnsSetPrivacyScopeError           = &ErrNo{Code: 10062, Message: "设置朋友圈可见范围失败"}
	SnsStrangerVisibilityEnabledError = &ErrNo{Code: 10063, Message: "设置陌生人可见失败"}
	SnsSetPrivacyError                = &ErrNo{Code: 10064, Message: "设置某条朋友圈权限失败"}
	SnsDeleteError                    = &ErrNo{Code: 10065, Message: "删除朋友圈失败"}
	SnsUploadVideoError               = &ErrNo{Code: 10066, Message: "上传视频失败"}
	SnsSendTextError                  = &ErrNo{Code: 10067, Message: "发送朋友圈文本失败"}
	SnsSendImageError                 = &ErrNo{Code: 10068, Message: "发送朋友圈图片失败"}
	SnsSendUrlError                   = &ErrNo{Code: 10069, Message: "发送朋友圈链接失败"}
	SnsSendVideoError                 = &ErrNo{Code: 10070, Message: "发送朋友圈视频失败"}
	SnsSendExcludeImageError          = &ErrNo{Code: 10071, Message: "转发非图片朋友圈失败"}
	SnsForwardError                   = &ErrNo{Code: 10072, Message: "转发朋友圈失败"}

	HearBeatError    = &ErrNo{Code: 10073, Message: "心跳失败"}
	SecAutoAuthError = &ErrNo{Code: 10074, Message: "自动授权失败"}
	LongOpenError    = &ErrNo{Code: 10075, Message: "长连接打开失败"}
)
