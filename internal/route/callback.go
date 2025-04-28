package router

import (
	"maogou/khan/internal/handler/callback"
	"maogou/khan/internal/handler/callback/chain"
	"maogou/khan/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initCallbackRoute(route *gin.Engine, sdk *khan.Khan) {
	verify := chain.NewVerify(sdk)
	ignore := chain.NewIgnore(sdk)
	pushMsg := chain.NewPushMsg(sdk)
	reconnection := chain.NewReconnection(sdk)
	redPacket := chain.NewRedPacket(sdk)
	payment := chain.NewPayment(sdk)

	verify.SetNext(ignore)
	ignore.SetNext(pushMsg)
	pushMsg.SetNext(reconnection)
	reconnection.SetNext(redPacket)
	redPacket.SetNext(payment)

	callbackHandler := callback.NewCallbackHandler(sdk, verify)

	route.POST("/api/v1/collect", callbackHandler.Collect)
	route.POST("/v1/api/monitor", callbackHandler.Monitor)
}
