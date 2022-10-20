package front

import (
	"gindemo01/public"
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

func Emailverification(c *gin.Context) {
	var E emailR
	E.text = "验证执行未知错误，请联系管理员"
	uid, ok1 := c.GetQuery("uid") //取不到query就返回false
	token, ok2 := c.GetQuery("token")
	emails, ok3 := c.GetQuery("email")
	times, ok4 := c.GetQuery("time")
	AntiModification, ok5 := c.GetQuery("antimodification")
	timeout, err := strconv.ParseInt(times, 10, 64)
	if !ok1 && !ok2 && !ok3 && !ok4 && ok5 && err != nil {
		E.text = "验证失败,链接缺少必要参数"
	} else if timeout < time.Now().Unix() {
		E.text = "链接已失效"
	} else if AntiModification != public.MD5(public.MD5(times+uid+token+emails)+emails+uid+emails) {
		E.text = "禁止修改链接"
	} else {
		U := public.UserInfoFind("email", emails, public.Method[0])
		if len(U) == 1 && U[0].Uid == uid && U[0].Token == token && U[0].Userstatus != 1 {
			U[0].Token = public.MD5(U[0].Token)
			U[0].Userstatus = 1
			public.UserInfoReviseStates(U[0])
			u := public.UserInfoFind("email", emails, public.Method[0])
			E.text = "验证成功，请前往首页重新登录"
			if u[0].Userstatus == 4 {
				E.text = "数据库写入失败请联系管理员"
			}
		} else if U[0].Userstatus == 1 {
			E.text = "已完成验证，请勿重复验证"
		} else {
			E.text = "验证失败，找不到此用户或校验识别码失效"
		}
	}
	S := time.Now().String()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "用户注册验证界面",
		"time":  S[0:16],
		"text":  E.text,
	})
}
