package login

import (
	"smallBot/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) Logout(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->Logout方法")
}
