package front

import (
	"gindemo01/public"
	"gindemo01/routes/front/Increased"
	"gindemo01/routes/front/corporate"
	"gindemo01/routes/front/identity"
	"gindemo01/routes/front/login"
	"gindemo01/routes/front/notice"
	"gindemo01/routes/front/personal"
	"gindemo01/routes/front/query"
	"gindemo01/routes/front/revise"
	"gindemo01/routes/front/token"
	"gindemo01/routes/front/upload"
	limits "github.com/gin-contrib/size"
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
		frontGroup.GET("/", index)                   //首页
		frontGroup.POST("/login", login.User)        //登录
		frontGroup.POST("/signup", login.SignUpUser) //注册
		frontGroup.GET("/email", emails)
		frontGroup.GET("/tokens", login.Setting)
		frontGroup.GET("/tokentime", login.Getting)
		frontGroup.GET("/query", query.QueryUserinfo)
		frontGroup.GET("/userinfo", token.Getting, query.UserInfo)
		frontGroup.PUT("/userinfo", token.Getting, revise.UserInfo)
		frontGroup.POST("/upload", limits.RequestSizeLimiter(4<<20), upload.Img)
		frontGroup.POST("/captcha", Increased.Captcha)
		frontGroup.POST("/personal", token.Getting, personal.PostPersonal)
		frontGroup.GET("/personal", token.Getting, personal.GetPersonal)
		frontGroup.POST("/corporate", token.Getting, corporate.PostCorporate)
		frontGroup.GET("/corporate", token.Getting, corporate.GetCorporate)
		frontGroup.GET("/identity", token.Getting, identity.Get)
		frontGroup.POST("/identity", token.Getting, identity.POST)
		frontGroup.PUT("/identity", token.Getting, identity.PUT)
		frontGroup.POST("/forgetpasswd", login.Forget)
		frontGroup.GET("/notice", notice.Get)
	}
}

// 前端主路由
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"state": 1,
		"msg":   "api在线中",
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
	public.Email(to, public.CaptchaEmail("?123"))
	c.JSON(200, gin.H{
		"code": 200,
	})

}
