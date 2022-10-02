package front

import (
	"fmt"
	"gindemo01/email"
	"gindemo01/sql_operate"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type JsonR struct {
	msg   int    //操作结果
	state int    //账号状态
	uid   string //账号id
	token string //识别码
}

func Main(rack *gin.Engine) {
	rack.GET("/", Index)
	frontGroup := rack.Group("/front")
	{
		frontGroup.GET("/", Index)
		frontGroup.POST("/login", signup)
		frontGroup.POST("/signup", SignUpUser)
		frontGroup.GET("/tokens", ResponseTokens)
		frontGroup.GET("/tokentime", CheckTokenTime)
		frontGroup.GET("/email", Email)
		frontGroup.GET("/signup/emailverification", emailverification)
	}
} // 前端主路由
func SignUpUser(c *gin.Context) {
	var u sql_operate.Usertoken
	u.Token = "error"
	tokenAdd := 1
	U1 := sql_operate.Userinfo{
		Uid:    c.PostForm("uid"),
		Name:   c.PostForm("username"),
		Passwd: universal.MD5(c.PostForm("passwd")),
		Email:  c.PostForm("email"),
	}
	U2 := sql_operate.Useremailtoken{
		Uid:            U1.Uid,
		Email:          U1.Email,
		Updatetime:     time.Now().Unix(),
		Expirationtime: time.Now().Unix() + 600,
		State:          0,
		Token:          universal.MD5(U1.Passwd + U1.Uid),
	}
	codes := 200
	states := 1
	if len(U1.Uid) != 11 {
		codes = 200
		states = 0 //表示错误来自客户端
	} else {
		states = sql_operate.UserInfoAdd(U1) //添加用户表单
		u.Uid = U1.Uid
		u.Token = universal.MD5(U1.Passwd + U1.Uid)
		states += sql_operate.UserTokenAdd(u) * 10
		//添加用户token表单
		states += sql_operate.UserEmailTokenAdd(U2) * 100
		//添加注册验证用表单
	}
	//fmt.Println(states)
	to := []string{U1.Email}
	add := "?uid=" + U1.Uid + "&token=" + u.Token + "&email=" + U1.Email
	if states > 111 {
		u.Token = "error"
	}
	JsonR := gin.H{
		"state": states, //表示状态
		"uid":   U1.Uid,
		"code":  codes,
		"token": u.Token,
		"msg":   tokenAdd,
		"email": email.Email(to, email.SignUp(add)),
	}
	c.JSON(200, JsonR)
}

// 用户注册

//	func LoginUser(c *gin.Context) {
//		u := sql_operate.Userinfo{
//			Uid:    c.PostForm("uid"),
//			Passwd: c.PostForm("passwd"),
//		}
//		U := sql_operate.Usertoken{
//			Uid: u.Uid,
//		}
//		tokens := "error"
//		u.Passwd = universal.MD5(u.Passwd)
//		var user = sql_operate.UserInfoFind(u)
//		states := 3
//		if len(user) == 0 {
//			states = 2
//		} else if user[0].Passwd == u.Passwd {
//			states = 1
//		}
//		if states == 1 {
//			tokens = sql_operate.UserTokenRevise(U)
//		}
//		times := time.Now().Unix()
//		JsonR := gin.H{
//			"time":  times,
//			"state": states,
//			"token": tokens,
//		}
//		c.JSON(200, JsonR)
//
// } //登录demo
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
func Email(c *gin.Context) {
	emails, ok1 := c.GetQuery("email")
	if !ok1 {
		c.JSON(200, gin.H{
			"error": "error",
		})
		return
	}
	to := []string{emails}
	email.Email(to, email.SignUp("?123"))
	c.JSON(200, gin.H{
		"code": 200,
	})

}
func emailverification(c *gin.Context) {
	uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	token, ok2 := c.GetQuery("token")
	emails, ok3 := c.GetQuery("email")
	U := sql_operate.Useremailtoken{
		Uid:   uid,
		Email: emails,
		Token: token,
	}
	U2 := sql_operate.Usertoken{
		Uid:   uid,
		Token: token,
	}
	if !ok3 && !ok1 && !ok2 {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "error",
		})
	} else {
		u := sql_operate.UserEmailTokenFind(U)
		u2 := sql_operate.UserTokenFind(U2)
		if len(u) != 1 {
			c.JSON(200, gin.H{
				"code": 201,
				"msg":  "error",
			})
		} else {
			c.JSON(200, gin.H{
				"code":  200,
				"msg":   "true",
				"token": sql_operate.UserTokenRevise(u2[0]),
			})
		}
	}
}
func signup(c *gin.Context) {
	var u sql_operate.Userinfo
	var R JsonR
	u.Uid = strconv.FormatInt(time.Now().Unix()-60000000, 10)
	u.Email = c.PostForm("email")
	u.Name = c.PostForm("username")
	u.Passwd = universal.MD5(c.PostForm("passwd") + time.Now().String())
	u.Token = universal.MD5(u.Passwd + u.Uid)
	u.Userstatus = 4
	u.Signtime = time.Now().String()
	R.uid = u.Uid
	R.state = u.Userstatus
	R.token = u.Token
	R.msg = sql_operate.UserInfoAdd(u)
	email.SignUp(u.Email)
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   R.msg,
		"state": R.state,
		"token": R.token,
		"uid":   R.uid,
	})

}
