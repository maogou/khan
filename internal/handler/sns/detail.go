package sns

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Detail(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Detail方法")

	var (
		req    v1.SnsDetailRequest
		result = v1.SnsDetailResponse{
			CommentList:  make([]v1.SnsCommentListItem, 0),
			LikeList:     make([]v1.SnsLikeListItem, 0),
			WithUserList: make([]v1.SnsWithUserListItem, 0),
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsDetail(
		ctx, sns.SnsDetailRequest{
			AppId:   req.AppId,
			Decrypt: true,
			Id:      req.SnsId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsDetail方法失败")
		response.Fail(ctx, errno.SnsDetailError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用SnsDetail方法失败")
		response.Fail(ctx, errno.SnsDetailError)
		return
	}

	result.Id = resp.Data.Object.Id
	result.UserName = resp.Data.Object.Username
	result.NickName = resp.Data.Object.Nickname
	result.CreateTime = resp.Data.Object.CreateTime
	result.SnsXml = resp.Data.Object.ObjectDesc.Buffer
	result.LikeCount = resp.Data.Object.LikeCount
	if len(resp.Data.Object.LikeUserList) > 0 {
		for _, v := range resp.Data.Object.LikeUserList {
			result.LikeList = append(
				result.LikeList, v1.SnsLikeListItem{
					CreateTime: v.CreateTime,
					NickName:   v.Nickname,
					Source:     v.Source,
					Type:       v.Type,
					UserName:   v.Username,
				},
			)
		}
	}
	result.CommentCount = resp.Data.Object.CommentCount

	if len(resp.Data.Object.CommentUserList) > 0 {
		for _, v := range resp.Data.Object.CommentUserList {
			result.CommentList = append(
				result.CommentList, v1.SnsCommentListItem{
					UserName:        v.Username,
					NickName:        v.Nickname,
					Source:          v.Source,
					Type:            v.Type,
					Content:         v.Content,
					CreateTime:      v.CreateTime,
					CommentId:       v.CommentId,
					ReplyCommentId:  v.ReplyCommentId,
					ReplyCommentId2: v.ReplyCommentId2,
					IsNotRichText:   v.IsNotRichText,
				},
			)
		}
	}
	result.WithUserCount = resp.Data.Object.WithUserCount

	if len(resp.Data.Object.WithUserList) > 0 {
		for _, v := range resp.Data.Object.WithUserList {
			result.WithUserList = append(
				result.WithUserList, v1.SnsWithUserListItem{
					Username:        v.Username,
					Source:          v.Source,
					Type:            v.Type,
					CreateTime:      v.CreateTime,
					CommentId:       v.CommentId,
					ReplyCommentId:  v.ReplyCommentId,
					IsNotRichText:   v.IsNotRichText,
					ReplyCommentId2: v.ReplyCommentId2,
					CommentFlag:     v.CommentFlag,
					CommentId2:      v.CommentId2,
					DeleteFlag:      v.DeleteFlag,
				},
			)
		}
	}

	response.Success(ctx, result)

}
