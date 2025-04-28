package contact

import (
	"encoding/json"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/contact"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
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

		cKey := "contact_cache_" + req.AppId

		cByte, err := json.Marshal(result)

		if err != nil {
			log.C(ctx).Error().Err(err).Msg("json序列化失败")
			response.Fail(ctx, errno.JsonEncodeError)
			return
		}

		if err = c.sdk.Rdb().Set(ctx, cKey, string(cByte), 10*time.Minute).Err(); err != nil {
			log.C(ctx).Error().Err(err).Msg("设置联系人缓存失败")
			response.Fail(ctx, errno.SetContactCacheError)
			return
		}

		response.Success(ctx, result)
		return
	case <-ctx.Done():
		log.C(ctx).Info().Msg("请求被取消")
		response.Fail(ctx, errno.RequestCanceled)
		return
	}
}
