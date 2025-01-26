package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/contact"
	"smallBot/internal/sdk/gewe"
)

func initContactRoute(engine *gin.Engine, sdk *gewe.Gewe) {
	contactHandler := contact.NewContactHandler(sdk)

	engine.POST("/v2/api/contacts/fetchContactsList", contactHandler.List)
	engine.POST("/v2/api/contacts/search", contactHandler.Search)
}
