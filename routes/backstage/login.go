package backstage

import (
	"gindemo01/public"
	"gindemo01/struct/sql_del_struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type loginR struct {
	Uid   string
	token string
	msg   int
	state int
}

func loginAdmin(c *gin.Context) {
	var u sql_del_struct.Admininfo
	var R loginR
	R.Uid = "?"
	R.token = "?"
	R.state = 0
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		U := public.AdmininfoFind("email = ?", u.Email)
		if len(U) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户不存在"})
		} else {
			R.state = U[0].State
			R.Uid = U[0].Uid
			if public.MD5(u.Passwd+R.Uid) == U[0].Passwd {
				R.token = SetTokenAdmininfo(U[0], time.Hour)
				c.JSON(200, gin.H{
					"msg":   R.msg,
					"state": R.state,
					"uid":   R.Uid,
					"token": R.token,
				})
			} else {
				c.JSON(200, gin.H{
					"msg": R.msg,
					"err": "密码错误",
				})
			}
		}
	}
}
