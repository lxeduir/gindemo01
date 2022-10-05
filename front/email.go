package front

import (
	"gindemo01/sql_operate"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type emailR struct {
	text     string
	username string
	time     string
}

func emailverification(c *gin.Context) {
	var E emailR
	E.text = "验证执行未知错误，请联系管理员"
	uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	token, ok2 := c.GetQuery("token")
	emails, ok3 := c.GetQuery("email")
	times, ok4 := c.GetQuery("time")
	timeout, err := strconv.ParseInt(times, 10, 64)
	if !ok1 && !ok2 && !ok3 && !ok4 && err != nil {
		E.text = "验证失败,链接缺少必要参数"
	} else if timeout < time.Now().Unix() {
		E.text = "验证失败,链接已失效"
	} else {
		U := sql_operate.UserInfoFindEmail(emails)
		if len(U) == 1 && U[0].Uid == uid && U[0].Token == token {
			U[0].Token = universal.MD5(U[0].Token)
			U[0].Userstatus = 1
			sql_operate.UserInfoRevise(U[0])
			E.text = "验证成功，请前往首页重新登录"
		} else if len(U) == 1 && U[0].Userstatus == 1 {
			E.text = "已完成验证，请勿重复验证"
		} else {
			E.text = "验证失败，找不到此用户或校验识别码失效"
		}

	}
	S := time.Now().String()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "用户注册验证界面",
		"time":  S[0:19],
		"text":  E.text,
	})
}
