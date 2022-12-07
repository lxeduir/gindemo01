package backstage

import (
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
)

func del(c *gin.Context) {
	uid, ok1 := c.GetQuery("uid")
	if !ok1 {
		c.JSON(201, gin.H{
			"msg": "缺少必须参数",
		})
	} else {
		m := sql.DelAdmininfo(uid)
		c.JSON(200, gin.H{
			"msg": m,
		})
	}
}
