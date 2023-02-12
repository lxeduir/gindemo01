package backstage

import (
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
)

func getPersonal(c *gin.Context) {
	cla, ok1 := c.Get("cla")
	Cla := cla.(Claimadmins)
	if !ok1 {
		c.JSON(201, gin.H{
			"err": "认证失败",
		})
	}
	affair, _ := sql.AffairsFind("affaries_id = ?", "%")
	c.JSON(200, gin.H{
		"list": affair,
		"cla":  Cla,
	})

}
