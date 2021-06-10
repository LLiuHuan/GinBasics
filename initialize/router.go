package initialize

import (
	"fairman-server/global"
	"fairman-server/middlewares"
	"fairman-server/router"
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
	global.FMS_LOG.Info("use middleware logger")
	// 跨域
	r.Use(middlewares.Cors())

	global.FMS_LOG.Info("use middleware cors")
	// swagger 文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	global.FMS_LOG.Info("register swagger handler")

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
