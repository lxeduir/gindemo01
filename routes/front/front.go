package front

import (
	"gindemo01/public"
	"gindemo01/routes/front/infofind"
	"gindemo01/routes/front/login"
	"github.com/gin-gonic/gin"
)

type Rjson struct {
	str  []string
	name string
}

func Main(rack *gin.Engine) {
	rack.GET("/", index)
	rack.LoadHTMLFiles("docs/html/index.html")
	frontGroup := rack.Group("/front")
	{
		frontGroup.GET("/", index)                                           //首页
		frontGroup.POST("/login", login.LoginUser)                           //登录
		frontGroup.POST("/signup", login.SignUpUser)                         //注册
		frontGroup.GET("/signup/emailverification", login.Emailverification) //邮件验证接口
		frontGroup.GET("/email", emails)
		frontGroup.GET("/tokens", login.Setting)
		frontGroup.GET("/tokentime", login.Getting)
		frontGroup.GET("/query", infofind.QueryUserinfo)
		frontGroup.GET("/userinfo", infofind.UserInfo)
	}
}

// 前端主路由
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"state": 1,
		"msg":   "api在线中",
		"code":  200,
	})
} //访问api根目录
func emails(c *gin.Context) {
	emails, ok1 := c.GetQuery("email")
	if !ok1 {
		c.JSON(200, gin.H{
			"msg": "error",
		})
		return
	}
	to := []string{emails}
	public.Email(to, public.SignUp("?123"))
	c.JSON(200, gin.H{
		"code": 200,
	})

}
