package contact

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
)

func (c *ContactHandler) ContactGetMobile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->ContactGetMobile方法")

	var (
		req    v1.ContactGetMobileRequest
		result = make([]v1.ContactGetMobileItem, 0)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactGetMobile(
		ctx, contact.ContactGetMobileRequest{
			Appid:   req.AppId,
			PhoneNo: lo.Ternary(len(req.Phones) > 0, req.Phones, make([]string, 0)),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->ContactGetMobile方法失败")
		response.Fail(ctx, errno.ContactGetMobileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用sdk->ContactGetMobile方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ContactGetMobileError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用sdk->ContactGetMobile方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.ContactGetMobileError)
		return
	}

	for _, item := range resp.Data.FriendList {
		result = append(
			result, v1.ContactGetMobileItem{
				UserName:        item.UserName,
				NickName:        item.Nickname,
				Sex:             item.Sex,
				PhoneMd5:        item.MobileMD5,
				Signature:       item.Signature,
				Country:         item.Country,
				Province:        item.Province,
				City:            item.City,
				BigHeadImgUrl:   item.BigHeadImgUrl,
				SmallHeadImgUrl: item.SmallHeadImgUrl,
				PersonalCard:    item.PersonalCard,
			},
		)
	}

	response.Success(ctx, result)

}
