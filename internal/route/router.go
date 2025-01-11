package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) error {
	initCommonRoute(engine)
	//initPprofRoute(engine)

	return nil
}
