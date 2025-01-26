package contact

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ContactHandler) Search(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->Search方法")

	var req v1.ContactSearchRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactSearch(
		ctx, transform.ContactSearchRequest{
			Appid:    req.AppId,
			UserName: req.ContactsInfo,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ContactSearch方法失败")
		response.Fail(ctx, errno.ContactSearchError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用ContactSearch方法失败")
		response.Fail(ctx, errno.ContactSearchError)
		return
	}

	response.Success(
		ctx, v1.ContactSearchResponse{
			V3:              resp.Data.UserName.String,
			NickName:        resp.Data.NickName.String,
			Sex:             resp.Data.Sex,
			Signature:       resp.Data.Signature,
			BigHeadImgUrl:   resp.Data.BigHeadImgUrl,
			SmallHeadImgUrl: resp.Data.SmallHeadImgUrl,
			V4:              "",
			Wxid:            resp.Data.Pyinitial.String,
		},
	)
}
