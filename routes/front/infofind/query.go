package infofind

import (
	"gindemo01/public"
	"github.com/gin-gonic/gin"
)

func QueryUserinfo(c *gin.Context) {
	method, ok1 := c.GetQuery("method")
	content, ok2 := c.GetQuery("content")
	if !ok1 && !ok2 {
		c.JSON(200, gin.H{
			"msg":  "参数错误",
			"code": 200,
		})
		return
	}
	var f public.Userinfo
	var find public.Finder = &f
	//U := public.UserInfoFind(method, content, manner)
	U := find.All(method, content).([]public.Userinfo)
	c.JSON(200, gin.H{
		"code": 200,
		"list": U,
	})
}
