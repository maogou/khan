package router

import (
	"smallBot/internal/handler/download"

	"github.com/gin-gonic/gin"
)

func initDownloadRouter(route *gin.Engine) {
	downloadHandler := download.NewDownloadHandler()

	route.GET("/api/v2/download/:filename", downloadHandler.StaticFile)
}
