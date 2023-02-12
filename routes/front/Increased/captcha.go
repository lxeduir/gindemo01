package Increased

import (
	"gindemo01/public"
	"gindemo01/public/redis"
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type reception struct {
	Email string
	Type  string
}

func Captcha(c *gin.Context) {
	var rec reception
	if err := c.ShouldBind(&rec); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "参数错误"})
	} else {
		var d int
		if rec.Type == "login" || rec.Type == "forget" {
			d = 3
		} else if rec.Type == "signup" {
			U := sql.UserinfoFind("email = ?", rec.Email)
			if len(U) != 0 {
				c.JSON(http.StatusOK, gin.H{"err": "邮箱已存在"})
				return
			}
			d = 4
		} else {
			c.JSON(http.StatusOK, gin.H{"err": "参数错误"})
			return
		}
		c2, _ := redis.Get(rec.Email, d)
		if c2 != "" {
			ic, _ := strconv.ParseInt(c2, 10, 64)
			t := 60 - (time.Now().Unix() - ic)
			c.JSON(http.StatusOK, gin.H{"err": "请" + strconv.FormatInt(t, 10) + "秒后再试"})
			return
		}
		caps := redis.SetCaptcha(rec.Email, time.Minute*5, d)
		err = redis.Set(rec.Email, strconv.FormatInt(time.Now().Unix(), 10), time.Minute, 3)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"err": "验证码获取失败"})
			return
		}
		to := []string{rec.Email}
		public.Email(to, public.CaptchaEmail(caps))
		c.JSON(http.StatusOK, gin.H{"msg": "验证码获取成功"})

	}
}
