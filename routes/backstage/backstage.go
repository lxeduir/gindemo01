package backstage

import (
	"github.com/gin-gonic/gin"
)

func BackMain(rback *gin.Engine) {
	BackGroup := rback.Group("/backstage")
	{
		BackGroup.POST("/login", loginAdmin)
		BackGroup.POST("/signup", signUpAdmin)
		BackGroup.GET("/admininfo", admininfo)
		BackGroup.GET("/del", del)
	}
} //登录demo
