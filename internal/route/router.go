package router

import (
	"maogou/khan/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, sdk *khan.Khan) error {
	initCommonRoute(engine)
	initLoginHandler(engine, sdk)
	initChatroomHandler(engine, sdk)
	initDownloadRouter(engine, sdk)
	initPersonal(engine, sdk)
	initCallbackRoute(engine, sdk)
	initMessageRoute(engine, sdk)
	initContactRoute(engine, sdk)
	initLabelRoute(engine, sdk)
	initSnsRoute(engine, sdk)
	initFavorRoute(engine, sdk)
	initTransferRoute(engine, sdk)
	initSettingRoute(engine, sdk)
	initPprofRoute(engine)

	return nil
}
