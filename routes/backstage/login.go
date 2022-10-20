package backstage

import (
	"gindemo01/config"
	"gindemo01/public"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoginAdmin(c *gin.Context) {
	var u config.Admininfo
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		var user = public.AdminInfoFind("email", u.Email, public.Method[0])
		states := 3
		if len(user) == 0 {
			states = 2
		}
		if user[0].Passwd == public.MD5(u.Passwd+user[0].Uid) {
			states = 1
		}
		tokens := user[0].Token
		times := time.Now().Unix()
		JsonR := gin.H{
			"uid":   user[0].Uid,
			"time":  times,
			"token": tokens,
			"state": states,
		}
		c.JSON(200, JsonR)
	}
}
