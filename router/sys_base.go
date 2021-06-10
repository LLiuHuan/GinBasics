package router

import (
	v1 "gin-basics/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("ping", v1.Ping)
	}
	return BaseRouter
}
