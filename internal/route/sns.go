package router

import (
	"smallBot/internal/handler/sns"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initSnsRoute(engine *gin.Engine, sdk *khan.Khan) {
	snsHandler := sns.NewSnsHandler(sdk)

	engine.POST("/v1/api/sns/uploadSnsImage", snsHandler.UploadImg)
	engine.POST("/v1/api/sns/snsDetails", snsHandler.Detail)
	engine.POST("/v1/api/sns/snsList", snsHandler.MyselfPage)
	engine.POST("/v1/api/sns/contactsSnsList", snsHandler.FriendPage)
	engine.POST("/v1/api/sns/likeSns", snsHandler.Like)
	engine.POST("/v1/api/sns/cancelLikeSns", snsHandler.CancelLike)
	engine.POST("/v1/api/sns/commentSns", snsHandler.Comment)
	engine.POST("/v1/api/sns/deleteCommentSns", snsHandler.DeleteComment)
	engine.POST("/v1/api/sns/snsVisibleScope", snsHandler.SetPrivacyScope)
	engine.POST("/v1/api/sns/strangerVisibilityEnabled", snsHandler.StrangerVisibility)
	engine.POST("/v1/api/sns/snsSetPrivacy", snsHandler.SetPrivacy)
	engine.POST("/v1/api/sns/delSns", snsHandler.Delete)
	engine.POST("/v1/api/sns/uploadSnsVideo", snsHandler.UploadVideo)
	engine.POST("/v1/api/sns/sendTextSns", snsHandler.SendText)
	engine.POST("/v1/api/sns/sendImageSns", snsHandler.SendImage)
	engine.POST("/v1/api/sns/sendUrlSns", snsHandler.SendUrl)
	engine.POST("/v1/api/sns/sendVideoSns", snsHandler.SendVideo)
	engine.POST("/v1/api/sns/forwardSns", snsHandler.Forward)
}
