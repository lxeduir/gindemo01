package revise

import (
	"fmt"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	var u sql_struct.Userinfo
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		fmt.Print(u)
	}
}
