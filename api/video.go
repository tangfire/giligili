package api

import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
)

func CreateVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{Code: 0, Msg: "成功"})
}
