package routes

import (
	"gindemo01/public"
	"gindemo01/routes/front"
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
		frontGroup.POST("/login", front.LoginUser)                           //登录
		frontGroup.POST("/signup", front.SignUpUser)                         //注册
		frontGroup.GET("/signup/emailverification", front.Emailverification) //邮件验证接口
		frontGroup.GET("/email", emails)
		frontGroup.GET("/tokens", front.Setting)
		frontGroup.GET("/tokentime", front.Getting)
		frontGroup.GET("/query", front.QueryUserinfo)
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
			"error": "error",
		})
		return
	}
	to := []string{emails}
	public.Email(to, public.SignUp("?123"))
	c.JSON(200, gin.H{
		"code": 200,
	})

}
