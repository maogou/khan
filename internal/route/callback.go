package router

import (
	"smallBot/internal/handler/callback"
	"smallBot/internal/handler/callback/chain"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initCallbackRoute(route *gin.Engine, sdk *gewe.Gewe) {
	ignore := chain.NewIgnore(sdk)
	pushMsg := chain.NewPushMsg(sdk)
	reconnection := chain.NewReconnection(sdk)

	ignore.SetNext(pushMsg)
	pushMsg.SetNext(reconnection)

	callbackHandler := callback.NewCallbackHandler(sdk, ignore)

	route.POST("/api/v1/collect", callbackHandler.Collect)
}
