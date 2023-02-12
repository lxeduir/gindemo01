package revise

import (
	"gindemo01/public/sql"
	"gindemo01/routes/front/token"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	var u sql_struct.Userinfo
	cla, _ := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		u.Uid = Cla.UserId
		err := sql.ReviseUserinfo(u)
		if err != nil {
			c.JSON(201, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "修改成功",
		})
	}
}
