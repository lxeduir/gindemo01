package login

import (
	"gindemo01/public"
	"gindemo01/public/redis"
	"gindemo01/public/sql"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var cnt int = 0

type Cjson struct {
	Username string
	Email    string
	Passwd   string
	Captcha  string
}

func SignUpUser(c *gin.Context) {
	var u sql_struct.Userinfo
	var cjson Cjson
	if err := c.ShouldBind(&cjson); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		u.Uid = strconv.FormatInt((time.Now().Unix()-660000000)*100+int64(r.Intn(128)), 10) //给定一个uid
		u.Passwd = public.MD5(cjson.Passwd + u.Uid)                                         //对密码进行加密
		caps, _ := redis.GetCaptcha(cjson.Captcha, 4)
		if caps == "err" {
			c.JSON(201, gin.H{"err": "验证码错误"})
			return
		} else {
			if caps == cjson.Email {
				U := sql.UserinfoFind("email = ?", cjson.Email)
				if len(U) > 0 {
					c.JSON(201, gin.H{"err": "该邮箱已经注册"})
					return
				}
				u.Email = cjson.Email
				u.Username = cjson.Username
				u.Permissions = "user"
				u.Userstatus = 1
				u.Signtime = time.Now().String()[0:19]
				sql.UserinfoAdd(u)
				c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})

			} else {
				c.JSON(201, gin.H{"err": "验证码错误"})
				return
			}
		}

	}
}
