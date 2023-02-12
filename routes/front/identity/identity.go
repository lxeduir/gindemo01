package identity

import (
	"gindemo01/public/sql"
	"gindemo01/routes/front/token"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	cla, ok1 := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	if ok1 != true {
		c.JSON(201, gin.H{"msg": "uid不能为空"})
		return
	}
	u := sql.UserIdentityFind("uid = ?", Cla.UserId)
	if len(u) == 0 {
		c.JSON(201, gin.H{
			"err": "无身份认证信息",
		})
		return
	}
	c.JSON(200, gin.H{
		"list": u,
	})
}
func POST(c *gin.Context) {
	var identity sql_struct.UserIdentity
	cla, ok1 := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	if ok1 != true {
		c.JSON(201, gin.H{"msg": "uid不能为空"})
		return
	}
	if err := c.ShouldBindJSON(&identity); err != nil {
		c.JSON(201, gin.H{
			"err": err,
		})
		return
	}
	identity.Uid = Cla.UserId
	err := sql.UserIdentityAdd(identity)
	if err != nil {
		c.JSON(201, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "添加成功",
	})

}
func PUT(c *gin.Context) {
	var identity sql_struct.UserIdentity
	cla, ok1 := c.Get("cla")
	Cla := cla.(token.Claimadmins)
	if ok1 != true {
		c.JSON(201, gin.H{"msg": "uid不能为空"})
		return
	}
	if err := c.ShouldBindJSON(&identity); err != nil {
		c.JSON(201, gin.H{
			"err": err.Error(),
		})
		return
	}
	identity.Uid = Cla.UserId
	err := sql.ReviseUserIdentity(identity)
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
