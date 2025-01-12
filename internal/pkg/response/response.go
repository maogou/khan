package response

import (
	"net/http"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

const SuccessCode = 0

type ErrResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data,omitempty"`
	Qid     string         `json:"qid"`
}

type ErrValidatorResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data,omitempty"`
	Qid     string            `json:"qid"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Qid     string      `json:"qid"`
}

type SuccessMsgResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Qid     string `json:"qid"`
}

// Fail 错误响应
func Fail(c *gin.Context, err error) {
	code, message := errno.DecodeErr(err)
	c.JSON(
		http.StatusOK, ErrResponse{
			Code:    code,
			Message: message,
			Qid:     c.GetString(constant.QID),
		},
	)
}

func ValidatorErr(c *gin.Context, err error, data map[string]string) {
	code, message := errno.DecodeErr(err)
	c.JSON(
		http.StatusOK, ErrValidatorResponse{
			Code:    code,
			Message: message,
			Data:    data,
			Qid:     c.GetString(constant.QID),
		},
	)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK, SuccessResponse{
			Code:    SuccessCode,
			Message: "success",
			Data:    data,
			Qid:     c.GetString(constant.QID),
		},
	)
}

func SuccessMsg(c *gin.Context, message string) {
	c.JSON(
		http.StatusOK, SuccessMsgResponse{
			Code:    SuccessCode,
			Message: message,
			Qid:     c.GetString(constant.QID),
		},
	)
}
