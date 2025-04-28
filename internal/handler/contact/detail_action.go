package contact

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/contact"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ContactHandler) Detail(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->Detail方法")

	var (
		req    v1.ContactDetailRequest
		result = make([]v1.ContactDetailItem, 0)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactDetail(
		ctx, contact.ContactDetailRequest{
			Appid:      req.AppId,
			ToWxidList: req.Wxids,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ContactDetail方法失败")
		response.Fail(ctx, errno.ContactDetailError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用ContactDetail方法失败")
		response.Fail(ctx, errno.ContactDetailError)
		return
	}

	for _, item := range resp.Data.ContactList {
		result = append(
			result, v1.ContactDetailItem{
				UserName:        item.UserName.String,
				NickName:        item.NickName.String,
				PyInitial:       item.Pyinitial.String,
				QuanPin:         item.QuanPin.String,
				Sex:             item.Sex,
				Remark:          item.Remark,
				RemarkPyInitial: item.RemarkPyinitial,
				RemarkQuanPin:   item.RemarkQuanPin,
				Signature:       item.Signature,
				Alias:           "",
				SnsBgImg:        item.SnsUserInfo.SnsBgimgId,
				Country:         item.Country,
				BigHeadImgUrl:   item.BigHeadImgUrl,
				SmallHeadImgUrl: item.SmallHeadImgUrl,
				Description:     "",
				CardImgUrl:      "",
				LabelList:       "",
				Province:        item.Province,
				City:            item.City,
				PhoneNumList:    "",
			},
		)
	}

	response.Success(ctx, result)
}
