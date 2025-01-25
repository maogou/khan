package router

import (
	"smallBot/internal/handler/download"

	"github.com/gin-gonic/gin"
)

func initDownloadRouter(route *gin.Engine) {
	downloadHandler := download.NewDownloadHandler()

	route.GET("/v2/api/download/:filename", downloadHandler.StaticFile)
}
