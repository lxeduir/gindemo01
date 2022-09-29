package front

import (
	"fmt"
	"gindemo01/sql_operate"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
	"time"
)

func Main(rack *gin.Engine) {
	rack.GET("/", Index)
	frontGroup := rack.Group("/front")
	{
		frontGroup.GET("/", Index)
		frontGroup.POST("/login", LoginUser)
		frontGroup.POST("/signup", SignUpUser)
		frontGroup.GET("/tokens", ResponseTokens)
		frontGroup.GET("/tokentime", CheckTokenTime)
	}
} //前端主路由
func SignUpUser(c *gin.Context) {
	var u sql_operate.Usertoken
	u.Token = "error"
	tokenAdd := 1
	U := sql_operate.Userinfo{
		Uid:    c.PostForm("uid"),
		Name:   c.PostForm("username"),
		Passwd: universal.MD5(c.PostForm("passwd")),
	}
	codes := 200
	states := 1
	if len(U.Uid) != 11 {
		codes = 200
		states = 2 //表示错误来自客户端
	} else {
		states = sql_operate.UserInfoAdd(U) //添加用户表单
		u.Uid = U.Uid
		u.Token = universal.MD5(U.Passwd + U.Uid)
		tokenAdd = sql_operate.UserTokenAdd(u)
		//添加用户token表单
	}
	//fmt.Println(states)
	JsonR := gin.H{
		"state": states, //表示状态
		"uid":   U.Uid,
		"code":  codes,
		"token": u.Token,
		"msg":   tokenAdd,
	}
	c.JSON(200, JsonR)

} //用户注册
func LoginUser(c *gin.Context) {
	u := sql_operate.Userinfo{
		Uid:    c.PostForm("uid"),
		Passwd: c.PostForm("passwd"),
	}
	u.Passwd = universal.MD5(u.Passwd)
	var user = sql_operate.UserInfoFind(u)
	states := 3
	if len(user) == 0 {
		states = 2
	}
	if user[0].Passwd == u.Passwd {
		states = 1
	}
	times := time.Now().Unix()
	JsonR := gin.H{
		"time":  times,
		"state": states,
	}
	c.JSON(200, JsonR)

} //登录demo
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"state": 1,
		"msg":   "200",
		"code":  200,
	})
} //访问api根目录
func ResponseTokens(c *gin.Context) {
	uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	if !ok1 {
		c.JSON(201, gin.H{
			"uid":   "error",
			"code":  "201",
			"token": "warning",
		})
		return
	}
	token, ok2 := c.GetQuery("token") //取不到query就返回false
	if !ok2 {
		c.JSON(201, gin.H{
			"uid":   uid,
			"code":  "201",
			"token": "error",
		})
		return
	}
	var U sql_operate.Usertoken
	U.Uid = uid
	u := sql_operate.UserTokenFind(U)
	if token != u[0].Token {
		c.JSON(201, gin.H{
			"uid":   uid,
			"code":  "201",
			"token": "error",
		})
		return
	}
	token = token + uid            //避免token出现重复
	U.Token = universal.MD5(token) //变化token
	U.Updatetime = time.Now().Unix()
	U.Expirationtime = U.Updatetime + 201
	sql_operate.UserTokenRevise(U)
	c.JSON(200, gin.H{
		"uid":   U.Uid,
		"token": U.Token,
		"msg":   true,
		"code":  200,
	})
} //刷新token
func CheckTokenTime(c *gin.Context) {
	state := 1
	msg := 0
	uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	if !ok1 {
		state = 2
	}
	token, ok2 := c.GetQuery("token") //取不到query就返回false
	if !ok2 {
		state = 3
	}
	var u sql_operate.Usertoken
	u.Token = token
	u.Uid = uid
	if sql_operate.UserTokenFindTime(u) {
		msg = 1
	}
	fmt.Println(msg)
	var Rjson = gin.H{
		"state": state,
		"msg":   msg,
	}
	c.JSON(200, Rjson)
}
