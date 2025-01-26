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
	engine.POST("/v2/api/contacts/addContacts", contactHandler.Add)
	engine.POST("/v2/api/contacts/deleteFriend", contactHandler.Delete)
	engine.POST("/v2/api/contacts/getBriefInfo", contactHandler.Info)
}
