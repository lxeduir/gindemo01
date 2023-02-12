package backstage

import (
	"gindemo01/public"
	"gindemo01/public/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Authentication(c *gin.Context) {
	cla, ok1 := c.Get("cla")
	Cla := cla.(Claimadmins)
	if ok1 != true {
		c.JSON(200, gin.H{"msg": "uid不能为空"})
		return
	} else {
		types := c.GetHeader("type")
		typeId, _ := strconv.Atoi(c.GetHeader("type_id"))
		operate, _ := strconv.Atoi(c.GetHeader("operate"))
		u := sql.AdmininfoFirst("uid", Cla.UserId)
		s, _ := sql.AdminRoleFind("role_id = ?", strconv.Itoa(u.RoleId))
		if len(s) == 0 {
			c.JSON(200, gin.H{
				"uid":     Cla.UserId,
				"role_id": u.RoleId,
				"err":     "找不到对应身份",
			})
		} else {
		}
		mp := public.Makemps(s[0].PermissionJson)
		if Permissionvalidation(mp, types, typeId, operate) {
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"uid":     Cla.UserId,
				"role_id": u.RoleId,
				"err":     "权限不足",
			})
		}
	}
	return
}

func Permissionvalidation(mp map[string]map[int]int, types string, typeId int, operate int) bool {
	if mp["0"][0] >= operate {
		return true
	}
	if mp[types][0] >= operate {
		return true
	}
	if mp[types][typeId] >= operate {
		return true
	} else {
		return false
	}
}
