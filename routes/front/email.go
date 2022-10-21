package front

import (
	"gindemo01/public"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type emailR struct {
	text     string
	username string
	time     string
}

func Emailverification(c *gin.Context) {
	var E emailR
	E.text = "验证执行未知错误，请联系管理员"
	uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	token, ok2 := c.GetQuery("token")
	emails, ok3 := c.GetQuery("email")
	AntiModification, ok5 := c.GetQuery("antimodification")
	public.GetToken(AntiModification)
	if !ok1 && !ok2 && !ok3 && !ok5 {
		E.text = "验证失败,链接缺少必要参数"
	} else if
	S := time.Now().String()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "用户注册验证界面",
		"time":  S[0:16],
		"text":  E.text,
	})
}
