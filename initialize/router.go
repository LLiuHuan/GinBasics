package initialize

import (
	"gin-basics/global"
	"gin-basics/middlewares"
	"gin-basics/router"
	"github.com/gin-gonic/gin"
	"net/http"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	var r = gin.Default()
	//r.Use(GinLogger(), GinRecovery(true))
	// Router.Use(middleware.LoadTls()  // https
	global.GB_LOG.Info("use middleware logger")
	// 跨域
	r.Use(middlewares.Cors())

	global.GB_LOG.Info("use middleware cors")
	// swagger 文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	global.GB_LOG.Info("register swagger handler")

	PublicGroup := r.Group("/v1")
	{
		router.InitBaseRouter(PublicGroup)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
