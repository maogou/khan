package router

import (
	"smallBot/internal/pkg/license"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, sdk *gewe.Gewe, l *license.License) error {
	initCommonRoute(engine)
	initDownloadRouter(engine)
	initCallbackRoute(engine, sdk, l)
	initMessageRoute(engine, sdk)
	initPprofRoute(engine)

	return nil
}
