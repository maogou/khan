package router

import (
	"smallBot/internal/handler/download"

	"github.com/gin-gonic/gin"
)

func initDownloadRouter(route *gin.Engine) {
	downloadHandler := download.NewDownloadHandler()

	route.GET("/v1/api/download", downloadHandler.StaticFile)
}
