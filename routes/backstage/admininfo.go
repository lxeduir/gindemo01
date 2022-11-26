package backstage

import (
	"gindemo01/public"
	"github.com/gin-gonic/gin"
)

func admininfo(c *gin.Context) {
	uid, ok1 := c.GetQuery("uid")
	if ok1 != true {
		c.JSON(200, gin.H{"code": 200, "msg": "uid不能为空"})
		return
	} else {
		u := public.AdmininfoFirst("uid", uid)
		c.JSON(200, gin.H{"code": 200, "Uid": uid, "admininfo": u})
		return
	}

}
