package front

import (
	"gindemo01/email"
	"gindemo01/sql_operate"
	"gindemo01/universal"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type SignUpUserR struct {
	msg   int    //操作结果
	state int    //账号状态
	uid   string //账号id
	token string //识别码
}

func SignUpUser(c *gin.Context) {
	var u sql_operate.Userinfo
	var R SignUpUserR
	u.Uid = strconv.FormatInt(time.Now().Unix()-660000000, 10)
	u.Email = c.PostForm("email")
	u.Name = c.PostForm("username")
	u.Passwd = universal.MD5(c.PostForm("passwd") + u.Uid)
	u.Token = universal.MD5(u.Passwd + u.Uid)
	u.Userstatus = 4
	S := time.Now().String()
	u.Signtime = S[0:19]
	R.uid = u.Uid
	R.state = u.Userstatus
	R.token = u.Token
	R.msg = sql_operate.UserInfoAdd(u)
	if R.msg == 2 {
		c.JSON(200, gin.H{
			"code":  200,
			"msg":   R.msg,
			"state": R.state,
		})
	} else {
		var U2 sql_operate.Usertoken
		U2.Uid = R.uid
		U2.Token = u.Token
		U2.Refreshtoken = universal.MD5(u.Token)
		U2.Rtepirationtime = time.Now().Unix() + 86400
		to := []string{u.Email}
		times := strconv.FormatInt(time.Now().Unix()+600, 10)
		AntiModification := universal.MD5(universal.MD5(times+u.Uid+u.Token+u.Email) + u.Email + u.Uid + u.Email)
		add := "?uid=" + u.Uid + "&token=" + u.Token + "&email=" + u.Email + "&time=" + times + "&antimodification=" + AntiModification
		sql_operate.UserTokenAdd(U2)
		email.Email(to, email.SignUp(add))
		c.JSON(200, gin.H{
			"code":  200,
			"msg":   R.msg,
			"state": R.state,
			"token": R.token,
			"uid":   R.uid,
		})
	}

}
