package task

import (
	"context"
	"maogou/khan/api/khan/v1/transform/login"
	"maogou/khan/api/khan/v1/transform/personal"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"

	"github.com/rs/xid"
	"github.com/sourcegraph/conc"
)

func (m *Monitor) forceLogout() {
	m.crontab.AddFunc(
		"@every 5m", func() {
			var (
				wg     conc.WaitGroup
				appIds []string
			)
			ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())
			log.C(ctx).Info().Msg("强制检测当前登录的微信号是否和当时许可证分配微信号是否一致")

			lic, err := help.License(ctx, m.sdk.Rdb())

			if err != nil {
				log.C(ctx).Error().Err(err).Msg("forceLogout=>获取许可证失败")
				return
			}

			p, err := help.Permission(ctx, m.sdk.Rdb())
			if err != nil {
				log.C(ctx).Error().Err(err).Msg("forceLogout=>获取license失败")
				return
			}

			appIds = p.AppId

			for _, appId := range appIds {
				wg.Go(
					func() {
						resp, pErr := m.sdk.PersonalProfile(
							ctx, personal.PersonalProfileRequest{
								AppId: appId,
							},
						)

						if pErr != nil {
							log.C(ctx).Error().Err(pErr).Str("appId", appId).Msg("获取个人资料失败")
							return
						}

						if resp.Ret != 0 {
							log.C(ctx).Error().Err(pErr).Str("appId", appId).Msg("resp.Ret != 0==>获取个人资料失败")
							return
						}

						wxid := resp.Data.UserInfo.UserName.String
						alias := resp.Data.UserInfo.Alias
						log.C(ctx).Info().Str("auth_wxid", lic.Cus).Str("wxid", wxid).Str(
							"alias", alias,
						).Msg("许可证授权分配的wxid")
						if wxid != lic.Cus && alias != lic.Sub {
							if _, lErr := m.sdk.Logout(ctx, login.LogoutRequest{Appid: appId}); lErr != nil {
								log.C(ctx).Error().Err(lErr).Str("appId", appId).Msg("强制退出登录失败")
								return
							}
							log.C(ctx).Info().Str("appId", appId).Msg("强制退出登录成功")

							return
						}

						log.C(ctx).Info().Str(
							"appid", appId,
						).Msg("检测当前登录的微信号是否和当时许可证分配微信号一致,不需要强制退出用户")

					},
				)
			}

			wg.Wait()
		},
	)
}
