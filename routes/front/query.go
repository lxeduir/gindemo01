package front

import (
	"gindemo01/public"
	"github.com/gin-gonic/gin"
)

func QueryUserinfo(c *gin.Context) {
	method, ok1 := c.GetQuery("method")
	content, ok2 := c.GetQuery("content")
	manner, ok3 := c.GetQuery("manner")
	if !ok1 && !ok2 && ok3 {
		c.JSON(200, gin.H{
			"msg":  "参数错误",
			"code": 200,
		})
		return
	}
	U := public.UserInfoFind(method, content, manner)
	c.JSON(200, gin.H{
		"code": 200,
		"list": U,
	})
}
