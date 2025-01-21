package errno

var (
	Ok                  = &ErrNo{Code: 0, Message: "OK"}
	PageNotFoundError   = &ErrNo{Code: 404, Message: "访问的页面不存在"}
	MethodNotAllowError = &ErrNo{Code: 405, Message: "访问的路由请求方式错误"}
	InternalServerError = &ErrNo{Code: 10001, Message: "服务器内部,请稍后再试"}
	ValidateError       = &ErrNo{Code: 10002, Message: "参数验证错误"}

	SendMsgError              = &ErrNo{Code: 10003, Message: "发送消息失败"}
	DownloadImgError          = &ErrNo{Code: 10004, Message: "下载图片失败"}
	DownloadNameEmptyError    = &ErrNo{Code: 10005, Message: "下载文件名不能为空"}
	DownloadFileNotExistError = &ErrNo{Code: 10006, Message: "下载文件不存在"}
	DownloadFileOpenError     = &ErrNo{Code: 10007, Message: "文件打开失败"}
	DownloadFileStatError     = &ErrNo{Code: 10008, Message: "获取文件信息失败"}
)
