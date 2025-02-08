package contact

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ContactHandler) Info(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->Info方法")

	var (
		req    v1.ContactInfoRequest
		result = make([]v1.ContactInfoItem, 0)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactInfo(
		ctx, contact.ContactInfoRequest{
			Appid:      req.AppId,
			ToWxidList: req.Wxids,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ContactInfo方法失败")
		response.Fail(ctx, errno.ContactInfoError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用ContactInfo方法失败")
		response.Fail(ctx, errno.ContactInfoError)
		return
	}

	for _, item := range resp.Data.ContactList {
		result = append(
			result, v1.ContactInfoItem{
				UserName:        item.Username,
				NickName:        item.Contact.NickName.String,
				PyInitial:       item.Contact.PyInitial.String,
				QuanPin:         item.Contact.QuanPin.String,
				Sex:             item.Contact.Sex,
				Remark:          item.Contact.Remark.String,
				RemarkPyInitial: item.Contact.RemarkPyinitial.String,
				RemarkQuanPin:   item.Contact.RemarkQuanPin.String,
				Signature:       "",
				Alias:           item.Contact.Alias,
				SnsBgImg:        "",
				Country:         item.Contact.Country,
				BigHeadImgUrl:   item.Contact.BigHeadImgUrl,
				SmallHeadImgUrl: item.Contact.SmallHeadImgUrl,
				Description:     "",
				CardImgUrl:      "",
				LabelList:       "",
				Province:        item.Contact.Province,
				City:            item.Contact.City,
				PhoneNumList:    "",
			},
		)
	}

	response.Success(ctx, result)

}
