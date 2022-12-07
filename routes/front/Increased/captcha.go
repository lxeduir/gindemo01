package Increased

import (
	"gindemo01/public"
	"gindemo01/public/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type reception struct {
	Email string
}

func Captcha(c *gin.Context) {
	var rec reception
	if err := c.ShouldBind(&rec); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "参数错误"})
	} else {
		c2, _ := redis.Get(rec.Email)
		if c2 != "" {
			ic, _ := strconv.ParseInt(c2, 10, 64)
			t := 60 - (time.Now().Unix() - ic)
			c.JSON(http.StatusOK, gin.H{"err": "请" + strconv.FormatInt(t, 10) + "秒后再试"})
			return
		}
		caps := redis.SetCaptcha(rec.Email)
		err = redis.Set(rec.Email, strconv.FormatInt(time.Now().Unix(), 10), time.Minute)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"err": "验证码获取失败"})
			return
		}
		to := []string{rec.Email}
		public.Email(to, public.SignUp(caps))
		c.JSON(http.StatusOK, gin.H{"msg": "验证码获取成功"})

	}
}
