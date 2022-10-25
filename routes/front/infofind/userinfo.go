package infofind

import (
	"gindemo01/public"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var u public.Userinfo
	var find public.Finder = &u
	uid, ok1 := c.GetQuery("uid")
	token := c.GetHeader("Authorization")
	if !ok1 {
		c.JSON(201, gin.H{
			"msg":  "uid?",
			"code": 201,
		})
	} else {
		msg := public.GetToken(token)
		if msg["msg"] == 1 {
			u = find.First("uid", uid).(public.Userinfo)
			u.Passwd = ""
			c.JSON(200, gin.H{
				"list": u,
				"msg":  msg["msg"],
				"code": 200,
			})
		} else {
			c.JSON(200, msg)
		}

	}
}
