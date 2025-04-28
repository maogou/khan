package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func (s *SnsHandler) SendImage(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->SendImage方法")

	var (
		req            v1.SnsSendImageRequest
		transformImage []sns.SnsSendImageItem
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	for _, v := range req.ImgInfos {
		transformImage = append(
			transformImage, sns.SnsSendImageItem{
				Id:        "0",
				Type:      "1",
				Url:       v.FileUrl,
				ThumbUrl:  v.ThumbUrl,
				Width:     strconv.Itoa(v.Width),
				Height:    strconv.Itoa(v.Height),
				TotalSize: strconv.Itoa(v.Length),
				Md5:       v.FileMd5,
			},
		)
	}

	resp, err := s.sdk.SnsSendImage(
		ctx, sns.SnsSendImageRequest{
			AppId:        req.AppId,
			Content:      req.Content,
			Media:        transformImage,
			AtUser:       lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			AllowUser:    lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			DisableUser:  lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			AllowTagId:   make([]string, 0),
			DisableTagId: make([]string, 0),
			Private:      req.Privacy,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->SendImage方法失败")
		response.Fail(ctx, errno.SnsSendImageError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->SendImage方法失败")
		response.Fail(ctx, errno.SnsSendImageError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用SnsHandler->SendImage方法失败")
		response.Fail(ctx, errno.SnsSendImageError)
		return
	}

	response.Success(
		ctx, v1.SnsSendImageResponse{
			Id:         resp.Data.SnsObject.Id,
			UserName:   resp.Data.SnsObject.Username,
			NickName:   resp.Data.SnsObject.Nickname,
			CreateTime: resp.Data.SnsObject.CreateTime,
		},
	)

}
