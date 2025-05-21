package serve

import (
	"context"
	"errors"
	"maogou/khan/internal/config"
	"maogou/khan/internal/middleware"
	router "maogou/khan/internal/route"
	"maogou/khan/internal/sdk/khan"
	"net/http"
	"os"
	"os/signal"

	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func run(conf config.Config, sdk *khan.Khan) error {
	gin.SetMode(conf.Mode)
	route := gin.Default()
	//route.Use(middleware.RequestId(), middleware.VerifyLicense(sdk.Rdb(), conf))
	route.Use(middleware.RequestId())
	addr := ":" + strconv.Itoa(conf.Port)

	if err := router.InitRouter(route, sdk); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: route,
	}

	log.Info().Msg("http-api服务访问地址==>http://127.0.0.1" + addr)
	log.Info().Msg("终止服务,请按键盘上 Ctrl+C 键退出服务")

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("监听" + addr + "端口失败")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.TimeOut)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("http-api服务异常,关闭失败")
		return err
	}

	log.Info().Msg("已终止http-api对外接口访问")

	return nil
}
