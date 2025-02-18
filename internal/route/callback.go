package router

import (
	"smallBot/internal/handler/callback"
	"smallBot/internal/handler/callback/chain"
	"smallBot/internal/pkg/license"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initCallbackRoute(route *gin.Engine, sdk *khan.Khan, l *license.License) {
	verify := chain.NewVerify(l)
	ignore := chain.NewIgnore(sdk)
	pushMsg := chain.NewPushMsg(sdk)
	reconnection := chain.NewReconnection(sdk)
	redPacket := chain.NewRedPacket(sdk)

	verify.SetNext(ignore)
	ignore.SetNext(pushMsg)
	pushMsg.SetNext(reconnection)
	reconnection.SetNext(redPacket)

	callbackHandler := callback.NewCallbackHandler(sdk, verify)

	route.POST("/api/v1/collect", callbackHandler.Collect)
}
