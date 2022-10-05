package front

import (
	"gindemo01/sql_operate"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
)

type LoginR struct {
	Uid   string
	token string
	msg   int
	state int
	code  int
}

func LoginUser(c *gin.Context) {
	var u sql_operate.Userinfo
	var R LoginR
	R.Uid = "?"
	R.token = "?"
	R.state = 0
	R.msg = 0
	R.code = 200
	u.Email = c.PostForm("email")
	Passwd := c.PostForm("passwd")
	U := sql_operate.UserInfoFindEmail(u.Email)
	if len(U) == 0 {
		c.JSON(200, gin.H{
			"code": R.code,
			"msg":  R.msg,
		})
	} else {
		R.state = U[0].Userstatus
		R.Uid = U[0].Uid
		if universal.MD5(Passwd+R.Uid) == U[0].Passwd {
			R.msg = 1
			c.JSON(200, gin.H{
				"code":  R.code,
				"msg":   R.msg,
				"state": R.state,
				"uid":   R.Uid,
				"token": R.token,
			})
		} else {
			c.JSON(200, gin.H{
				"code":   R.code,
				"msg":    R.msg,
				"passwd": 0,
			})
		}

	}
}
