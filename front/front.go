package front

import (
	"fmt"
	"gindemo01/email"
	"gindemo01/sql_operate"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
	"time"
)

func Main(rack *gin.Engine) {
	rack.GET("/", Index)
	rack.LoadHTMLFiles("html/index.html")
	frontGroup := rack.Group("/front")
	{
		frontGroup.GET("/", Index)
		frontGroup.POST("/login", LoginUser)
		frontGroup.POST("/signup", SignUpUser)
		frontGroup.GET("/signup/emailverification", emailverification)
		frontGroup.GET("/tokens", ResponseTokens)
		frontGroup.GET("/tokentime", CheckTokenTime)
		frontGroup.GET("/email", Email)

	}
} // 前端主路由
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
