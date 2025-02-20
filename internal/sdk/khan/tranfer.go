package khan

import (
	"context"
	"smallBot/api/khan/v1/transform/tranfer"
	"smallBot/internal/pkg/log"
)

func (k *Khan) BatchHeartBeat(ctx context.Context, req tranfer.TranferBatchHeartBeatRequest) (*tranfer.TranferBatchHeartBeatResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&tranfer.TranferBatchHeartBeatResponse{}).Post(batchHeartBeat)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用BatchHeartBeat方法失败")
		return nil, err
	}

	return resp.Result().(*tranfer.TranferBatchHeartBeatResponse), nil
}
