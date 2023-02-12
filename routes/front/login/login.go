package login

import (
	"gindemo01/public"
	"gindemo01/public/redis"
	"gindemo01/public/sql"
	"gindemo01/routes/front/token"
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
type login struct {
	Email   string
	Passwd  string
	Captcha string
}

func User(c *gin.Context) {
	var R loginR
	var l login
	R.Uid = "?"
	R.token = "?"
	R.state = 0
	R.msg = 0
	if err := c.ShouldBind(&l); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		U := sql.UserinfoFind("email = ?", l.Email)
		if len(U) == 0 {
			c.JSON(201, gin.H{
				"err": "邮箱不存在",
			})
		} else {
			emails, _ := redis.GetCaptcha(l.Captcha, 3)
			if U[0].Passwd == public.MD5(l.Passwd+U[0].Uid) {
				R.state = U[0].Userstatus
				R.Uid = U[0].Uid
				R.token = token.SetTokenUserinfo(U[0], time.Hour)
				R.msg = 1
				c.JSON(200, gin.H{
					"msg":      R.msg,
					"uid":      R.Uid,
					"token":    R.token,
					"state":    R.state,
					"username": U[0].Username,
				})
			} else {

				if emails == U[0].Email {
					R.state = U[0].Userstatus
					R.Uid = U[0].Uid
					R.token = token.SetTokenUserinfo(U[0], time.Hour)
					R.msg = 1
					c.JSON(200, gin.H{
						"msg":      R.msg,
						"uid":      R.Uid,
						"token":    R.token,
						"state":    R.state,
						"username": U[0].Username,
					})
					return
				} else if l.Passwd == "" {
					c.JSON(201, gin.H{
						"err": "验证码错误",
					})
					return
				}

				c.JSON(201, gin.H{
					"err": "密码错误",
				})
				return
			}

		}
	} //登录函数
}
