package backstage

import (
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

func admininfo(c *gin.Context) {
	cla, ok1 := c.Get("cla")
	Cla := cla.(Claimadmins)
	if ok1 != true {
		c.JSON(200, gin.H{"msg": "uid不能为空"})
		return
	} else {
		u := sql.AdmininfoFind("uid LIKE ?", Cla.UserId)
		u[0].Passwd = ""
		c.JSON(200, gin.H{"admininfo": u})
		return
	}

}
func userinfo(c *gin.Context) {
	cla, ok1 := c.Get("admininfo")
	Cla := cla.(Claimadmins)
	if ok1 != true {
		c.JSON(200, gin.H{"msg": "uid不能为空"})
		return
	} else {
		typeId, _ := strconv.Atoi(c.Param("type_id"))
		operate, _ := strconv.Atoi(c.Param("operate"))
		if Permissionvalidation(Cla.Mps, "userinfo", typeId, operate) {
			u := sql.UserinfoFind("uid LIKE ?", "%")
			c.JSON(200, gin.H{
				"userinfo": u,
			})
			return
		} else {
			c.JSON(200, gin.H{
				"err": "权限不足",
			})
		}
	}
}
