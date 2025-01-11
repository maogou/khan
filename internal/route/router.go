package router

import (
	"smallBot/internal/sdk/gewe"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine, sdk *gewe.Gewe) error {
	initCommonRoute(engine)
	//initPprofRoute(engine)

	log.Info().Msgf("sdk=>%#v", sdk)

	return nil
}
