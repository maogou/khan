package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// initPprofRoute  注册pprof路由
func initPprofRoute(engine *gin.Engine) {
	pprof.Register(engine, "debug/pprof")
}
