package backstage

import (
	"gindemo01/public"
	"github.com/gin-gonic/gin"
)

func admininfo(c *gin.Context) {
	uid, ok1 := c.GetQuery("uid")
	token := c.GetHeader("Authorization")
	if !ok1 {
		c.JSON(201, gin.H{
			"msg":  "uid?",
			"code": 201,
		})
	} else {
		msg := public.GetTokenAdmin(token)
		if msg["msg"] == 1 {
			//A = query.First("uid", uid).(sql_struct.Admininfo)
			A := public.AdmininfoFirst("uid", uid)
			A.Passwd = ""
			c.JSON(200, gin.H{
				"data": gin.H{
					"uid":      A.Uid,
					"email":    A.Email,
					"username": A.Username,
					"state":    A.State,
				},
				"msg":  msg["msg"],
				"code": 200,
			})
		} else {
			c.JSON(200, msg)
		}

	}

}
