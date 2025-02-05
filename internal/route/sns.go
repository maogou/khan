package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/sns"
	"smallBot/internal/sdk/khan"
)

func initSnsRoute(engine *gin.Engine, sdk *khan.Khan) {
	snsHandler := sns.NewSnsHandler(sdk)

	engine.POST("/v2/api/sns/uploadSnsImage", snsHandler.UploadImg)
}
