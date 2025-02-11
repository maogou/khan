package router

import (
	"smallBot/internal/handler/sns"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initSnsRoute(engine *gin.Engine, sdk *khan.Khan) {
	snsHandler := sns.NewSnsHandler(sdk)

	engine.POST("/v2/api/sns/uploadSnsImage", snsHandler.UploadImg)
	engine.POST("/v2/api/sns/snsDetails", snsHandler.Detail)
	engine.POST("/v2/api/sns/snsList", snsHandler.MyselfPage)
	engine.POST("/v2/api/sns/contactsSnsList", snsHandler.FriendPage)
	engine.POST("/v2/api/sns/likeSns", snsHandler.Like)
	engine.POST("/v2/api/sns/cancelLikeSns", snsHandler.CancelLike)
	engine.POST("/v2/api/sns/commentSns", snsHandler.Comment)
	engine.POST("/v2/api/sns/deleteCommentSns", snsHandler.DeleteComment)
	engine.POST("/v2/api/sns/snsVisibleScope", snsHandler.SetPrivacyScope)
	engine.POST("/v2/api/sns/strangerVisibilityEnabled", snsHandler.StrangerVisibility)
	engine.POST("/v2/api/sns/snsSetPrivacy", snsHandler.SetPrivacy)
	engine.POST("/v2/api/sns/delSns", snsHandler.Delete)
	engine.POST("/v2/api/sns/uploadSnsVideo", snsHandler.UploadVideo)
	engine.POST("/v2/api/sns/sendTextSns", snsHandler.SendText)
	engine.POST("/v2/api/sns/sendImageSns", snsHandler.SendImage)
	engine.POST("/v2/api/sns/sendUrlSns", snsHandler.SendUrl)
}
