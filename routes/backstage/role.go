package backstage

import (
	"gindemo01/public"
	"gindemo01/struct/sql_del_struct"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetRole(c *gin.Context) {
	role := public.AdminRoleFind()
	c.JSON(200, gin.H{
		"list": role,
	})
}
func PostRole(c *gin.Context) {
	var u sql_struct.AdminRole
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		u.Ctime = time.Now().String()[0:19]
		u.Utime = u.Ctime
		msg := public.AdminRoleAdd(u)
		role := public.AdminRoleFind()
		c.JSON(200, gin.H{
			"msg":  msg,
			"list": role,
		})
	}

}
func PutRole(c *gin.Context) {
	var u sql_struct.AdminRole
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		msg := public.ReviseAdminrole(u)
		role := public.AdminRoleFind()
		c.JSON(200, gin.H{
			"msg":  msg,
			"list": role,
		})

	}
}
func DelRole(c *gin.Context) {
	var u sql_struct.AdminRole
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		var U sql_del_struct.AdminRole
		U.RoleId = u.RoleId
		msg := public.DelAdminRole(U)
		role := public.AdminRoleFind()
		c.JSON(200, gin.H{
			"msg":  msg,
			"list": role,
		})
	}
}
