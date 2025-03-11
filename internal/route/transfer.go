package router

import (
	"smallBot/internal/handler/transfer"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initTransferRoute(engine *gin.Engine, skd *khan.Khan) {
	transferHandler := transfer.NewTransferHandler(skd)

	engine.GET("/v1/api/transfer/backup", transferHandler.Backup)
	engine.POST("/v1/api/transfer/import", transferHandler.Import)
}
