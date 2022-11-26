package Increased

import (
	"fmt"
	"gindemo01/public"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"time"
)

type reception struct {
	Email string
}

func Captcha(c *gin.Context) {
	var captcha sql_struct.UserRedis
	var R reception
	if err := c.ShouldBind(&R); err != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
	} else {
		u := public.UserinfoFind("email", R.Email)
		if len(u) == 0 {
			c.JSON(400, gin.H{
				"err": "邮箱不存在",
			})
		} else {
			captcha.Uid = u[0].Uid
			captcha.Data = public.Captcha()
			captcha.Type = "验证码"
			captcha.CreateTime = time.Now().String()[0:19]
			captcha.ExpirationTime = time.Now().Add(time.Minute * 5).String()[0:19]
			public.UserRedisAdd(captcha)
			c.JSON(200, gin.H{
				"code": 200,
				"list": captcha,
			})
		}
	}
	captchaExpirationTime()
}
func captchaExpirationTime() {
	u := public.UserRedisFind("type", "验证码")
	for _, v := range u {
		formatTime, err := time.Parse("2006-01-02T15:04:05+08:00", v.ExpirationTime)
		if err != nil {
			fmt.Println(err)
		} else {
			if formatTime.Unix()-time.Now().Unix()-28800 < 0 {
				public.DelUserRedis(v)
			}

		}
	}
}
