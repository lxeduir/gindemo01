package front

import (
	"gindemo01/config"
	"gindemo01/public"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginR struct {
	Uid   string
	token string
	msg   int
	state int
	code  int
}

func LoginUser(c *gin.Context) {
	var u config.Userinfo
	var R loginR
	R.Uid = "?"
	R.token = "?"
	R.state = 0
	R.msg = 0
	R.code = 200
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		U := public.UserInfoFind("email", u.Email, public.Method[0])
		if len(U) == 0 {
			c.JSON(201, gin.H{
				"code": R.code,
				"msg":  R.msg,
			})
		} else {
			R.state = U[0].Userstatus
			R.Uid = U[0].Uid
			if public.MD5(u.Passwd+R.Uid) == U[0].Passwd {
				if R.state == 1 {
					R.token = SetToken(R.Uid)
					R.msg = 1
				} else if R.state == 4 {
					R.msg = 2
					R.Uid = "?"
				}

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
} //登录函数
