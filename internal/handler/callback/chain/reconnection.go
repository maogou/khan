package chain

import (
	"context"
	"errors"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"

	"github.com/spf13/cast"
)

type Reconnection struct {
	BaseHandler
	sdk *khan.Khan
}

func (r *Reconnection) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.C(ctx).Info().Msg("调用Reconnection->HandlerRequest方法")

	if r.IsCanHandler(ctx, param) {
		if err := r.Process(ctx, param); err != nil {
			log.C(ctx).Error().Err(err).Msg("调用Reconnection->Process方法失败")
		}
	}

}

func (r *Reconnection) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用Reconnection->IsCanHandler方法")

	msgType := param.TypeName

	switch msgType {
	case constant.TypeNameSystemMessages:
		data := cast.ToString(param.Data)
		if data != "长链服务已关闭" {
			return true
		}
	case constant.TypeNameSystemSuccess:
		return false
	case constant.TypeNameSystemClose, constant.TypeNameSystemFailure, constant.TypeNameSystemTimeout:
		return true
	case constant.TypeNameMsgAddMsg, constant.TypeNameMsgDelContacts, constant.TypeNameMsgModContacts, constant.TypeNameMsgFinderMsg:
		return false
	default:
		return false
	}

	return false
}

func (r *Reconnection) Process(ctx context.Context, param v1.CollectRequest) error {
	log.C(ctx).Info().Msg("调用Reconnection->Process方法")

	hbResp, err := r.sdk.HearBeat(
		ctx, v1.HearBeatRequest{
			AppId: r.sdk.GetAppId(),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用心跳包失败")
		return err
	}

	if hbResp.Ret != 0 {
		log.C(ctx).Warn().Str("err_msg", hbResp.MsgErr).Msg("调用心跳包失败")
		return errors.New(hbResp.MsgErr)
	}

	log.C(ctx).Info().Msg("调用心跳包成功")

	saaResp, err := r.sdk.SecAutoAuth(
		ctx, v1.SecAutoAuthRequest{
			AppId:  r.sdk.GetAppId(),
			SdkVer: "8.0.48",
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SecAutoAuth方法失败")
		return err
	}

	if saaResp.Ret != 0 {
		log.C(ctx).Warn().Str("err_msg", saaResp.MsgErr).Msg("调用SecAutoAuth方法失败")
		return errors.New(saaResp.MsgErr)
	}

	log.C(ctx).Info().Msg("调用SecAutoAuth方法成功")

	loResp, err := r.sdk.LongOpen(
		ctx, v1.LongOpenRequest{
			AppId:      r.sdk.GetAppId(),
			CleanCache: false,
			Host:       r.sdk.Config().Sdk.Collect,
			Timeout:    r.sdk.Config().TimeOut,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用开启长连接失败")
		return err
	}

	if loResp.Ret != 0 {
		log.C(ctx).Warn().Str("err_msg", loResp.MsgErr).Msg("调用开启长连接失败")
		return errors.New(loResp.MsgErr)
	}

	return nil

}

func NewReconnection(sdk *khan.Khan) *Reconnection {
	return &Reconnection{
		sdk: sdk,
	}
}

var _ Chain = (*Reconnection)(nil)
