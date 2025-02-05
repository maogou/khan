package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/contact"
	"smallBot/internal/sdk/khan"
)

func initContactRoute(engine *gin.Engine, sdk *khan.Khan) {
	contactHandler := contact.NewContactHandler(sdk)

	engine.POST("/v2/api/contacts/fetchContactsList", contactHandler.List)
	engine.POST("/v2/api/contacts/search", contactHandler.Search)
	engine.POST("/v2/api/contacts/addCont˚acts", contactHandler.Add)
	engine.POST("/v2/api/contacts/deleteFriend", contactHandler.Delete)
	engine.POST("/v2/api/contacts/getBriefInfo", contactHandler.Info)
	engine.POST("/v2/api/contacts/getDetailInfo", contactHandler.Detail)
	engine.POST("/v2/api/contacts/setFriendPermissions", contactHandler.OnlyChat)
	engine.POST("/v2/api/contacts/setFriendRemark", contactHandler.Remark)
}
