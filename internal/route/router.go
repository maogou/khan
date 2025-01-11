package router

import (
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, sdk *gewe.Gewe) error {
	initCommonRoute(engine)
	//initPprofRoute(engine)

	return nil
}
