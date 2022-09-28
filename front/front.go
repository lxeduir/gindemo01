package front

import (
	"gindemo01/sql_operate"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "服务器在线",
	})
} //api主页请求
func Main(rack *gin.Engine) {
	rack.GET("/", IndexHtml)
	frontGroup := rack.Group("/front")
	{
		frontGroup.POST("/login", LoginUser)
		frontGroup.POST("/signup", SignUpUser)
	}
}
func SignUpUser(c *gin.Context) {
	U := sql_operate.Userinfo{
		Uid:    c.PostForm("uid"),
		Name:   c.PostForm("username"),
		Passwd: sql_operate.MD5(c.PostForm("passwd")),
		Token:  sql_operate.MD5(c.PostForm("token")),
	}
	codes := 200
	states := 1
	if len(U.Uid) != 11 {
		codes = 200
		states = 2 //表示错误来自客户端
	} else {
		states = sql_operate.UserInfoAdd(U)
	}
	//fmt.Println(states)
	JsonR := gin.H{
		"state": states, //表示状态
		"uid":   U.Uid,
		"code":  codes,
	}
	c.JSON(200, JsonR)

} //用户注册
func LoginUser(c *gin.Context) {
	u := sql_operate.Userinfo{
		Uid:    c.PostForm("uid"),
		Passwd: c.PostForm("passwd"),
	}
	u.Passwd = sql_operate.MD5(u.Passwd)
	var user = sql_operate.UserInfoFind(u)
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
