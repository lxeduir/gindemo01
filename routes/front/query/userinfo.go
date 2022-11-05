package query

import (
	"gindemo01/public"
	"gindemo01/struct/sql_del_struct"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var u sql_del_struct.Userinfo
	uid, ok1 := c.GetQuery("uid")
	token := c.GetHeader("Authorization")
	if !ok1 {
		c.JSON(201, gin.H{
			"msg":  "uid?",
			"code": 201,
		})
	} else {
		msg := public.GetTokenUser(token)
		if msg["msg"] == 1 {

			u = public.UserinfoFirst("uid", uid)
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
