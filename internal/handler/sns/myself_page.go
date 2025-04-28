package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) MyselfPage(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->MyselfPage方法")

	var (
		req    v1.SnsMyselfPageRequest
		result = v1.SnsMyselfPageResponse{
			SnsList: make([]v1.SnsPageItem, 0),
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsMyselfPage(
		ctx, sns.SnsMyselfPageRequest{
			AppId:        req.AppId,
			Decrypt:      req.Decrypt,
			FirstPageMd5: req.FirstPageMd5,
			MaxId:        req.MaxId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsMyselfPage方法失败")
		response.Fail(ctx, errno.SnsMyselfPageError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用SnsMyselfPage方法失败")
		response.Fail(ctx, errno.SnsMyselfPageError)
		return
	}

	firstPageMd5 := resp.Data.FirstPageMd5
	objectCount := resp.Data.ObjectCount

	if objectCount > 0 && len(firstPageMd5) > 0 && firstPageMd5 != req.FirstPageMd5 {
		result.FirstPageMd5 = firstPageMd5
	}

	result.RequestTime = resp.Data.NewRequestTime
	result.SnsCount = objectCount

	if len(resp.Data.ObjectList) > 0 {
		result.MaxId = resp.Data.ObjectList[len(resp.Data.ObjectList)-1].Id

		for _, v := range resp.Data.ObjectList {
			likes := make([]v1.SnsLikeListItem, 0)
			for _, like := range v.LikeUserList {
				likes = append(
					likes, v1.SnsLikeListItem{
						UserName:   like.Username,
						NickName:   like.Nickname,
						Source:     like.Source,
						Type:       like.Type,
						CreateTime: like.CreateTime,
					},
				)
			}

			comments := make([]v1.SnsCommentListItem, 0)
			for _, comment := range v.CommentUserList {
				comments = append(
					comments, v1.SnsCommentListItem{
						UserName:        comment.Username,
						NickName:        comment.Nickname,
						Source:          comment.Source,
						Type:            comment.Type,
						Content:         comment.Content,
						CreateTime:      comment.CreateTime,
						CommentId:       comment.CommentId,
						ReplyCommentId:  comment.ReplyCommentId,
						ReplyCommentId2: comment.ReplyCommentId2,
						IsNotRichText:   comment.IsNotRichText,
					},
				)
			}

			withUsers := make([]v1.SnsWithUserListItem, 0)
			for _, withUser := range v.WithUserList {
				withUsers = append(
					withUsers, v1.SnsWithUserListItem{
						Username:        withUser.Username,
						Source:          withUser.Source,
						Type:            withUser.Type,
						CreateTime:      withUser.CreateTime,
						CommentId:       withUser.CommentId,
						ReplyCommentId:  withUser.ReplyCommentId,
						IsNotRichText:   withUser.IsNotRichText,
						ReplyCommentId2: withUser.ReplyCommentId2,
						CommentId2:      withUser.CommentId2,
						DeleteFlag:      withUser.DeleteFlag,
					},
				)
			}

			result.SnsList = append(
				result.SnsList, v1.SnsPageItem{
					Id:            v.Id,
					UserName:      v.Username,
					NickName:      v.Nickname,
					CreateTime:    v.CreateTime,
					SnsXml:        v.ObjectDesc.Buffer,
					LikeCount:     v.LikeCount,
					LikeList:      likes,
					CommentCount:  v.CommentCount,
					CommentList:   comments,
					WithUserCount: v.CommentUserListCount,
					WithUserList:  withUsers,
				},
			)
		}
	}

	response.Success(ctx, result)

}
