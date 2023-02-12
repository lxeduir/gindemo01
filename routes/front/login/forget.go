package login

import (
	"gindemo01/public"
	"gindemo01/public/redis"
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
)

type forget struct {
	Email   string
	Passwd  string
	Captcha string
}

func Forget(c *gin.Context) {
	var f forget
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(201, gin.H{
			"err": err.Error(),
		})
		return
	}
	email, err := redis.GetCaptcha(f.Captcha, 3)
	if err != nil {
		c.JSON(201, gin.H{
			"err": err.Error(),
		})
	}
	if email == f.Email {
		u, errs := sql.UserinfoFirst("email", f.Email)
		if errs != nil {
			c.JSON(201, gin.H{
				"err": errs.Error(),
			})
			return
		}
		err := sql.ReviseUserPasswd(f.Email, public.MD5(f.Passwd+u.Uid))
		if err != nil {
			c.JSON(201, gin.H{
				"err": "验证码有误或过期",
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "修改成功",
			})
		}
	}

}
