package router

import (
	"smallBot/internal/pkg/license"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, sdk *khan.Khan, l *license.License) error {
	initCommonRoute(engine)
	initLoginHandler(engine, sdk)
	initChatroomHandler(engine, sdk)
	initDownloadRouter(engine)
	initPersonal(engine, sdk)
	initCallbackRoute(engine, sdk, l)
	initMessageRoute(engine, sdk)
	initContactRoute(engine, sdk)
	initLabelRoute(engine, sdk)
	//initSnsRoute(engine, sdk)
	initFavorRoute(engine, sdk)
	initPprofRoute(engine)

	return nil
}
