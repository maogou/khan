package contact

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"
)

func (c *ContactHandler) List(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->List方法")

	var (
		req    v1.ContactListRequest
		sg     singleflight.Group
		result = v1.ContactListResponse{
			Chatrooms: make([]string, 0),
			Friends:   make([]string, 0),
			Ghs:       make([]string, 0),
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	// 使用 DoChan 方法
	ch := sg.DoChan(
		req.AppId, func() (interface{}, error) {
			var (
				currentChatRoomContactSeq int = 0
				currentWxcontactSeq       int = 0
				allResults                []*contact.ContactListResponse
			)

			for {
				resp, err := c.sdk.ContactList(
					ctx, contact.ContactListRequest{
						Appid:                     req.AppId,
						CurrentChatRoomContactSeq: currentChatRoomContactSeq,
						CurrentWxcontactSeq:       currentWxcontactSeq,
					},
				)
				if err != nil {
					return nil, err
				}

				if resp.Ret != 0 {
					return nil, errors.New("获取联系人列表失败")
				}

				allResults = append(allResults, resp)

				if resp.Data.CountinueFlag == 0 {
					break
				}

				// 更新分页信息
				currentChatRoomContactSeq = resp.Data.CurrentChatRoomContactSeq
				currentWxcontactSeq = resp.Data.CurrentWxcontactSeq
			}

			return allResults, nil
		},
	)

	// 等待结果
	select {
	case res := <-ch:
		if res.Err != nil {
			log.C(ctx).Error().Err(res.Err).Msg("获取联系人列表失败")
			response.Fail(ctx, errno.ContactListError)
			return
		}
		contactList := res.Val.([]*contact.ContactListResponse)

		for _, contact := range contactList {
			if len(contact.Data.ContactUsernameList) > 0 {

				for _, friend := range contact.Data.ContactUsernameList {
					if len(friend) > 0 && strings.HasPrefix(friend, "gh_") {
						result.Ghs = append(result.Ghs, friend)
					} else if len(friend) > 0 && strings.HasPrefix(friend, "wxid_") {
						result.Friends = append(result.Friends, friend)
					} else if len(friend) > 0 {
						result.Chatrooms = append(result.Chatrooms, friend)
					}
				}
			}
		}

		response.Success(ctx, result)
		return
	case <-ctx.Done():
		log.C(ctx).Info().Msg("请求被取消")
		response.Fail(ctx, errno.RequestCanceled)
		return
	}
}
