package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) Create(ctx *gin.Context) {
	log.C(ctx).Info().Msg("ChatRoomHandler->Create方法")

	var req v1.ChatRoomCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->Create方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomCreate(
		ctx, chatroom.ChatroomCreateRequest{
			Appid:      req.AppId,
			ToWxidList: req.Wxids,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomCreate方法失败")
		response.Fail(ctx, errno.ChatroomCreateError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用ChatroomCreate方法失败")
		response.Fail(ctx, errno.ChatroomCreateError)
		return
	}

	response.Success(
		ctx, v1.ChatRoomCreateResponse{
			ChatroomId:    resp.Data.ChatRoomName.String,
			HeadImgBase64: "data:image/jpg;base64," + resp.Data.ImgBuf.Buffer,
		},
	)

}
