package router

import (
	"maogou/khan/internal/handler/contact"
	"maogou/khan/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initContactRoute(engine *gin.Engine, sdk *khan.Khan) {
	contactHandler := contact.NewContactHandler(sdk)

	engine.POST("/v1/api/contacts/fetchContactsList", contactHandler.List)
	engine.POST("/v1/api/contacts/fetchContactsListCache", contactHandler.FetchContactsListCache)
	engine.POST("/v1/api/contacts/search", contactHandler.Search)
	engine.POST("/v1/api/contacts/addContacts", contactHandler.Add)
	engine.POST("/v1/api/contacts/deleteFriend", contactHandler.Delete)
	engine.POST("/v1/api/contacts/getBriefInfo", contactHandler.Info)
	engine.POST("/v1/api/contacts/getDetailInfo", contactHandler.Detail)
	engine.POST("/v1/api/contacts/setFriendPermissions", contactHandler.OnlyChat)
	engine.POST("/v1/api/contacts/setFriendRemark", contactHandler.Remark)
	engine.POST("v1/api/contacts/uploadPhoneAddressList", contactHandler.ContactUploadMobile)
	engine.POST("v1/api/contacts/getPhoneAddressList", contactHandler.ContactGetMobile)
}
