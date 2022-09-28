package backstage

import (
	"gindemo01/sql_operate"
	"github.com/gin-gonic/gin"
	"time"
)

func Main(rback *gin.Engine) {
	BackGroup := rback.Group("/backstage")
	{
		BackGroup.POST("/login", LoginAdmin)
		BackGroup.POST("/signup", SignUpUser)
	}
}
func SignUpUser(c *gin.Context) {
	U := sql_operate.Admininfo{
		Uid:    c.PostForm("uid"),
		Name:   c.PostForm("username"),
		Passwd: sql_operate.MD5(c.PostForm("passwd")),
		Token:  sql_operate.MD5(c.PostForm("token")),
		Rtoken: sql_operate.MD5(c.PostForm("rtoken")),
	}
	codes := 200
	states := 1
	if len(U.Uid) != 11 {
		codes = 200
		states = 2 //表示错误来自客户端
	} else {
		states = sql_operate.AdminInfoAdd(U)
	}
	//fmt.Println(states)
	JsonR := gin.H{
		"state": states, //表示状态
		"uid":   U.Uid,
		"code":  codes,
	}
	c.JSON(200, JsonR)
} //注册
func LoginAdmin(c *gin.Context) {
	u := sql_operate.Admininfo{
		Uid:    c.PostForm("uid"),
		Passwd: c.PostForm("passwd"),
	}
	u.Passwd = sql_operate.MD5(u.Passwd)
	var user = sql_operate.AdminInfoFind(u)
	states := 3
	if len(user) == 0 {
		states = 2
	}
	if user[0].Passwd == u.Passwd {
		states = 1
	}
	tokens := user[0].Token
	times := time.Now().Unix()
	JsonR := gin.H{
		"time":  times,
		"token": tokens,
		"state": states,
	}
	c.JSON(200, JsonR)

} //登录demo
