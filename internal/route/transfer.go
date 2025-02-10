package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/transfer"
	"smallBot/internal/sdk/khan"
)

func initTransferRoute(engine *gin.Engine, skd *khan.Khan) {
	transferHandler := transfer.NewTransferHandler(skd)

	engine.GET("/v2/api/transfer/backup", transferHandler.Backup)
	engine.POST("/v2/api/transfer/import", transferHandler.Import)
}
