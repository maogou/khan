package router

import (
	"smallBot/internal/handler/sns"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initSnsRoute(engine *gin.Engine, sdk *khan.Khan) {
	snsHandler := sns.NewSnsHandler(sdk)

	engine.POST("/v2/api/sns/uploadSnsImage", snsHandler.UploadImg)
	engine.POST("/v2/api/sns/snsDetails", snsHandler.Detail)
}
