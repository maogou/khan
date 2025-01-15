package router

import (
	"smallBot/internal/pkg/license"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, sdk *gewe.Gewe, l *license.License) error {
	initCommonRoute(engine)
	initCallbackRoute(engine, sdk, l)
	//initPprofRoute(engine)

	return nil
}
