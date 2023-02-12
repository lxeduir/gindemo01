package backstage

import (
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
)

func deladmin(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(Claimadmins)
	if Cla.RoleId == "1" {
		uid, ok1 := c.GetQuery("uid")
		if !ok1 {
			c.JSON(201, gin.H{
				"msg": "缺少必须参数",
			})
		} else {
			sql.DelAdmininfo(uid)
			u := sql.AdmininfoFind("uid = ?", uid)
			if len(u) == 0 {
				c.JSON(200, gin.H{
					"msg": "删除成功",
				})
			} else {
				c.JSON(201, gin.H{
					"msg": "删除失败",
				})
			}
		}
	} else {
		c.JSON(201, gin.H{
			"msg": "权限不足",
		})
	}

}
func deluser(c *gin.Context) {
	cla, _ := c.Get("cla")
	Cla := cla.(Claimadmins)
	if Cla.RoleId == "1" {
		uid, ok1 := c.GetQuery("uid")
		if !ok1 {
			c.JSON(201, gin.H{
				"msg": "缺少必须参数",
			})
		} else {
			sql.DelUserinfo(uid)
			u := sql.UserinfoFind("uid = ?", uid)
			if len(u) == 0 {
				c.JSON(200, gin.H{
					"msg": "删除成功",
				})
			} else {
				c.JSON(201, gin.H{
					"msg": "删除失败",
				})
			}
		}
	} else {
		c.JSON(201, gin.H{
			"msg": "权限不足",
		})
	}

}
