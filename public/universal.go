package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorApi(rack *gin.Engine) {
	rack.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":   "来到了荒漠",
			"code":  200,
			"api":   "error",
			"error": "请检测api路径是否正确",
		})
	})
}
