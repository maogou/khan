package help

import (
	"context"
	"smallBot/internal/constant"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

func GetQidLog(ctx context.Context) zerolog.Logger {
	if zl, ok := ctx.Value(constant.ZEROLOG).(zerolog.Logger); ok {
		return zl
	}

	if qid, ok := ctx.Value(constant.QID).(string); ok {
		return log.With().Str(constant.QID, qid).Logger()
	}

	return log.Logger
}
