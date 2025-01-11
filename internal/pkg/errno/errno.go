package errno

import "fmt"

type ErrNo struct {
	Code    int
	Message string
}

func (err ErrNo) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func NewErr(errNo *ErrNo, err error) *Err {
	return &Err{
		Code:    errNo.Code,
		Message: errNo.Message,
		Err:     err,
	}
}

// Add 修改错误信息
func (err *Err) Add(message string) error {
	err.Message += " " + message

	return err
}

// AddFormat 使用格式化函数修改错误信息
func (err *Err) AddFormat(format string, args ...any) error {
	err.Message += " " + fmt.Sprintf(format, args...)

	return err
}

// Error 实现error接口
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr 解析错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Ok.Code, Ok.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *ErrNo:
		return typed.Code, typed.Message
	default:
		return InternalServerError.Code, InternalServerError.Message
	}
}
