package login

import (
	"github.com/gin-gonic/gin"
)

type emailR struct {
	text     string
	username string
	time     string
}

func Emailverification(c *gin.Context) {
	//var E emailR
	//E.text = "验证执行未知错误，请联系管理员"
	//uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	//token, ok2 := c.GetQuery("token")
	//emails, ok3 := c.GetQuery("email")
	//AntiModification, ok5 := c.GetQuery("antimodification")
	//u := public.GetTokenUser(AntiModification)
	//if u["msg"] == "token不能为空" {
	//	E.text = "token不能为空"
	//} else if u["msg"] == "token错误" {
	//	E.text = "token错误"
	//} else if u["msg"] == "用户不存在" {
	//	E.text = "用户不存在"
	//}
	//if !ok1 && !ok2 && !ok3 && !ok5 {
	//	E.text = "验证失败,链接缺少必要参数"
	//} else {
	//	S := time.Now().String()
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"Title": "用户注册验证界面",
	//		"time":  S[0:16],
	//		"text":  E.text,
	//	})
	//}
}
