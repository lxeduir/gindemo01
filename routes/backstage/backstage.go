package backstage

import (
	"gindemo01/routes/front/infofind"

	"github.com/gin-gonic/gin"
)

func BackMain(rback *gin.Engine) {
	BackGroup := rback.Group("/backstage")
	{
		BackGroup.POST("/login", LoginAdmin)
		BackGroup.POST("/signup", SignUpAdmin)
		BackGroup.GET("/userinfo", infofind.UserInfo)
	}
} //登录demo
