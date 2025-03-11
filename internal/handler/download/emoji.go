package download

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/download"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (d *DownloadHandler) DownloadEmoji(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用DownloadHandler->Emoji方法")

	var req v1.DownloadEmojiRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->Emoji方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := d.sdk.DownloadEmoji(
		ctx, download.DownloadEmojiRequest{
			Appid:   req.AppId,
			Md5List: []string{req.EmojiMd5},
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->Emoji方法调用失败")
		response.Fail(ctx, errno.DownloadEmojiError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->DownloadHandler->Emoji方法调用失败")
		response.Fail(ctx, errno.DownloadEmojiError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->DownloadHandler->Emoji方法调用失败")
		response.Fail(ctx, errno.DownloadEmojiError)
		return
	}

	if len(resp.Data.EmojiList) == 0 {
		log.C(ctx).Error().Msg("len(resp.Data.EmojiList) == 0 ->DownloadHandler->Emoji方法调用失败")
		response.Fail(ctx, errno.DownloadEmojiError)
		return
	}

	response.Success(
		ctx, v1.DownloadEmojiResponse{
			Url: resp.Data.EmojiList[0].Url,
		},
	)
}
