package backstage

import (
	"github.com/gin-gonic/gin"
)

func BackMain(rback *gin.Engine) {
	BackGroup := rback.Group("/backstage")
	{
		BackGroup.POST("/login", loginAdmin)
		BackGroup.POST("/signup", signUpAdmin)
		BackGroup.GET("/admininfo", Getting, admininfo)
		BackGroup.GET("/del", del)
		BackGroup.GET("/path", Getting, path)
		BackGroup.GET("/role", GetRole)
		BackGroup.POST("/role", PostRole)
		BackGroup.PUT("/role", PutRole)
		BackGroup.DELETE("/role", DelRole)

	}
} //登录demo
