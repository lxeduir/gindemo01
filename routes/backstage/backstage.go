package backstage

import (
	"github.com/gin-gonic/gin"
)

func BackMain(rback *gin.Engine) {
	BackGroup := rback.Group("/backstage")
	{
		BackGroup.POST("/login", loginAdmin)
		BackGroup.POST("/signup", signUpAdmin)
		BackGroup.GET("/admininfo", getting, admininfo)
		BackGroup.GET("/del", del)
		BackGroup.GET("/path", getting, path)
		BackGroup.GET("/role", getting, GetRole)
		BackGroup.POST("/role", getting, PostRole)
		BackGroup.PUT("/role", getting, PutRole)
		BackGroup.DELETE("/role", getting, DelRole)
		BackGroup.GET("/permission", GetPermission)
		BackGroup.POST("/permission", PostPermission)
		BackGroup.PUT("/permission", PutPermission)
		BackGroup.DELETE("/permission", DelPermission)
		BackGroup.POST("/authentication", getting, Authentication)
		BackGroup.GET("/personal", getting, getPersonal)
	}
	AuthGroup := BackGroup.Group("/auth")
	{
		AuthGroup.POST("/userinfo/:type_id/:operate", getting, userinfo)

	}
}

func Auth(c *gin.Context) {
	cla, ok1 := c.Get("admininfo")
	if ok1 != true {
		c.JSON(200, gin.H{"code": 200, "msg": "uid不能为空"})
		return
	} else {
		types := c.Param("type")
		typeId := c.Param("type_id")
		operate := c.Param("operate")
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  types + typeId + operate,
			"cla":  cla,
		})
		return
	}

}
