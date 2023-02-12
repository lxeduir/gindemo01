package query

import (
	"gindemo01/public/sql"
	"gindemo01/routes/front/token"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	u := sql.UserinfoFind("uid = ?", Cla.UserId)
	if len(u) == 0 {
		c.JSON(201, gin.H{
			"err": "无用户信息",
		})
		return
	}
	u[0].Passwd = ""
	c.JSON(200, u[0])
	return
}
