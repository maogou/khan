package login

import (
	"context"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/login"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/sdk/khan"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"

	"github.com/rs/xid"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pterm/pterm"
	"github.com/skip2/go-qrcode"

	"github.com/rs/zerolog/log"
)

type QuickLogin struct {
	template
	tableData  pterm.TableData
	url        string
	sdk        *khan.Khan
	wxid       string
	p          *tea.Program
	selected   int
	appId      string
	zLog       zerolog.Logger
	ctx        context.Context
	permission *v1.Permission
	err        error
}

var _ BootStrap = (*QuickLogin)(nil)

func NewQuickLogin(sdk *khan.Khan, p *tea.Program) *QuickLogin {
	ql := new(QuickLogin)
	qid := xid.New().String()
	ql.ctx = context.WithValue(context.Background(), "qid", qid)
	ql.sdk = sdk
	ql.p = p
	ql.tableData = pterm.TableData{{"配置项", "对应值", "备注"}}
	ql.b = ql
	ql.zLog = log.With().Str("qid", qid).Logger()

	return ql
}

func (q *QuickLogin) SetAppId(appId string) {
	p, err := help.Permission(q.ctx, q.sdk.Rdb())
	if err != nil {
		q.err = err
	} else if !slices.Contains(p.AppId, q.appId) {
		q.err = errors.New("appId未授权")
	} else {
		q.sdk.SetAppId(appId)
		q.appId = appId
		q.permission = p
	}
}

func (q *QuickLogin) GetLoginQrCode() {
	lqcResp, err := q.sdk.LoginQrCode(q.ctx, login.GetLoginQrCodeRequest{AppId: q.appId})
	if err != nil {
		q.err = err
		q.zLog.Error().Err(err).Msg("获取二维码失败")
	} else if lqcResp.Ret != 0 {
		q.zLog.Warn().Str("err_msg", lqcResp.Msg).Msg("调用获取二维码失败")
		q.err = errors.New(lqcResp.Msg)
	} else {
		q.tableData = append(
			q.tableData,
			[]string{"uuid", lqcResp.Data.Uuid, "登录随机生成的唯一标识"},
			[]string{"nkey", lqcResp.Data.Nkey, "钥匙串key"},
			[]string{"appId", q.appId, "设备标识id(很重要)"},
			[]string{"welcome", "欢迎使用khan服务!", "文件传输助手欢迎消息"},
		)

		q.sdk.SetUuId(lqcResp.Data.Uuid)
		q.sdk.SetNKey(lqcResp.Data.Nkey)
		q.url = lqcResp.Data.Url

		cKey := constant.WXQrCodeCache + q.appId

		if err = q.sdk.Rdb().Set(
			q.ctx, cKey, lqcResp.Data.Uuid, constant.WXQrCodeExpire,
		).Err(); err != nil {
			q.zLog.Error().Err(err).Msg("设置wx的qrcode失败")
			q.err = errno.SetWxLoginQrCodeCacheError
		}
	}
}

func (q *QuickLogin) PrintQrCode() {
	qrCodeUrl := "http://weixin.qq.com/x/" + q.sdk.GetUuId()
	if qr, err := qrcode.New(qrCodeUrl, qrcode.Medium); err != nil {
		q.err = err
		q.zLog.Error().Err(err).Msg("生成二维码失败")
	} else {
		art := qr.ToSmallString(false)
		pterm.Println("请使用微信扫一扫登录")
		pterm.Println(art)
	}
}

func (q *QuickLogin) Confirm() {
	if m, err := q.p.Run(); err != nil {
		q.err = err
		q.zLog.Error().Err(err).Msg("运行确认登录失败")
	} else {
		if cm, ok := m.(confirm); ok && cm.selected != "" {
			for i := 0; i < len(cm.choices); i++ {
				if cm.choices[i] == cm.selected {
					q.selected = i + 1
					break
				}
			}

			if q.selected != 1 {
				q.err = errors.New("你取消了登录")
			}
		} else {
			q.err = errors.New("你退出确认登录")
		}
	}

}

func (q *QuickLogin) CheckLogin() {
	q.zLog.Info().Msg("正在检查登录状态,请稍等...")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	maxRetry := 11
	retryCount := 0

	for range ticker.C {
		q.zLog.Info().Msg("正在检查登录状态,请稍等...")
		retryCount++

		if retryCount > maxRetry {
			pterm.Println("登录超时")
			q.err = errors.New("登录超时")
			break
		}

		resp, cErr := q.sdk.CheckLoginQrCode(
			q.ctx, login.CheckLoginRequest{
				Appid: q.appId,
				Nkey:  q.sdk.GetNKey(),
				Uuid:  q.sdk.GetUuId(),
			},
		)

		if cErr != nil {
			q.zLog.Error().Err(cErr).Msg("检查登录状态失败")
			continue
		}

		if resp.Ret != 0 {
			q.zLog.Warn().Str("err_msg", resp.Msg).Msg("调用检查登录状态失败")
			continue
		}
		q.wxid = resp.Data.LoginInfo.AcctSectResp.UserName
		lcKey := constant.WXLoginCache + q.appId
		if err := q.sdk.Rdb().Set(q.ctx, lcKey, q.wxid, 0).Err(); err != nil {
			q.zLog.Error().Err(err).Msg("设置wx登录缓存失败")
			q.err = errno.SetWxLoginCacheError
		} else {
			q.zLog.Info().Msg("亲爱的 (" + resp.Data.StatusInfo.NickName + ") 恭喜你成功登录!")
		}
		break

	}

}

func (q *QuickLogin) PrintConfig() {
	q.zLog.Info().Msg("确认登录后,请保存以下服务启动需要的配置如下:")

	q.err = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(q.tableData).Render()
}

func (q *QuickLogin) HearBeat() {
	q.zLog.Info().Msg("正在发送心跳包,请稍等...")
	hbResp, err := q.sdk.HearBeat(q.ctx, v1.HearBeatRequest{AppId: q.appId})

	if err != nil {
		q.zLog.Error().Err(err).Msg("发送心跳包失败")
		q.err = err
	} else {
		if hbResp.Ret != 0 {
			q.zLog.Warn().Str("err_msg", hbResp.MsgErr).Msg("调用发送心跳包失败")
			q.err = errors.New(hbResp.MsgErr)
		} else {
			q.zLog.Info().Msg("发送心跳包成功")
		}
	}
}

func (q *QuickLogin) AutoAuth() {
	q.zLog.Info().Msg("正在自动认证,请稍等...")
	saaResp, err := q.sdk.SecAutoAuth(q.ctx, v1.SecAutoAuthRequest{AppId: q.appId})

	if err != nil {
		q.zLog.Error().Err(err).Msg("自动认证失败")
		q.err = err
	} else {
		if saaResp.Ret != 0 {
			q.zLog.Warn().Str("err_msg", saaResp.MsgErr).Msg("调用自动认证失败")
			q.err = errors.New(saaResp.MsgErr)
		} else {
			q.zLog.Info().Msg("自动认证成功")
		}
	}
}

func (q *QuickLogin) Open() {
	q.zLog.Info().Msg("正在打开长连接,请稍等...")

	port := strconv.Itoa(q.sdk.Config().Port)

	loResp, err := q.sdk.LongOpen(
		q.ctx, v1.LongOpenRequest{
			AppId:      q.appId,
			CleanCache: true,
			Host:       "http://127.0.0.1:" + port + "/api/v1/collect",
			Timeout:    60,
		},
	)

	if err != nil {
		q.zLog.Error().Msg("调用开启长连接失败")
		q.err = err
	} else {
		if loResp.Ret != 0 {
			q.zLog.Warn().Str("err_msg", loResp.MsgErr).Msg("调用开启长连接失败")
			q.err = errors.New(loResp.MsgErr)
		} else {
			q.zLog.Info().Msg("开启长连接成功")
		}
	}
}

func (q *QuickLogin) Welcome() {
	//var welcome = message.PostTextRequest{
	//	AppId: q.appId,
	//	ToWxidList: []message.ToWxid{
	//		{
	//			Content: "欢迎使用Khan服务!",
	//			ToWxid:  "filehelper",
	//			MsgType: 1,
	//		},
	//	},
	//}
	//
	//var resp *message.PostTextResponse

	//resp, q.err = q.sdk.PostText(q.ctx, welcome)

	//q.zLog.Info().Any("welcome", resp).Msg("发送欢迎消息")

}

func (q *QuickLogin) CanNext() bool {
	return q.err == nil || q.selected == 1
}
