package v1

import (
	"gin-basics/model/response"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.OkWithData("pong", c)
}
