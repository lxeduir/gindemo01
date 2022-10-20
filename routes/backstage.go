package routes

import (
	"gindemo01/routes/backstage"
	"github.com/gin-gonic/gin"
)

func BackMain(rback *gin.Engine) {
	BackGroup := rback.Group("/backstage")
	{
		BackGroup.POST("/login", backstage.LoginAdmin)
		BackGroup.POST("/signup", backstage.SignUpAdmin)
	}
} //登录demo
