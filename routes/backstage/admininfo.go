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
	} else if Cla.RoleId != "1" {
		u := sql.AdmininfoFind("uid = ?", Cla.UserId)
		u[0].Passwd = ""
		c.JSON(200, gin.H{
			"uid":      u[0].Uid,
			"email":    u[0].Email,
			"username": u[0].Username,
			"state":    u[0].State,
			"roleid":   u[0].RoleId,
		})
		return
	} else {
		u := sql.AdmininfoFind("uid LIKE ?", "%")
		for i := 0; i < len(u); i++ {
			u[i].Passwd = ""
			u[i].Token = ""
		}
		c.JSON(200, gin.H{
			"admininfo": u,
		})
	}

}
func userinfo(c *gin.Context) {
	cla, ok1 := c.Get("cla")
	Cla := cla.(Claimadmins)
	if ok1 != true {
		c.JSON(200, gin.H{"msg": "uid不能为空"})
		return
	} else {
		typeId, _ := strconv.Atoi(c.Param("type_id"))
		operate, _ := strconv.Atoi(c.Param("operate"))
		if Permissionvalidation(Cla.Mps, "userinfo", typeId, operate) {
			u := sql.UserinfoFind("uid LIKE ?", "%")
			for i := 0; i < len(u); i++ {
				u[i].Passwd = ""
			}
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
