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

	return log.Logger
}
