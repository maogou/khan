package router

import (
	"smallBot/internal/handler/download"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initDownloadRouter(route *gin.Engine, sdk *khan.Khan) {
	downloadHandler := download.NewDownloadHandler(sdk)

	route.GET("/v1/api/download", downloadHandler.StaticFile)
	route.POST("/v1/api/download/emoji", downloadHandler.DownloadEmoji)
	route.POST("/v1/api/download/cdn", downloadHandler.DownloadCdn)
	route.POST("/v1/api/download/video", downloadHandler.DownloadVideo)
	route.POST("/v1/api/download/file", downloadHandler.DownloadFile)
}
