package login

import (
	"gindemo01/public"
	"gindemo01/struct/sql_struct"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var cnt int = 0

type SignUpUserR struct {
	msg   int    //操作结果
	state int    //账号状态
	uid   string //账号id
	token string //识别码
}

func SignUpUser(c *gin.Context) {
	var u sql_struct.Userinfo
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	} else {
		var R SignUpUserR
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		u.Uid = strconv.FormatInt((time.Now().Unix()-660000000)*100+int64(r.Intn(128)), 10) //给定一个uid
		u.Passwd = public.MD5(u.Passwd + u.Uid)                                             //对密码进行加密
		u.Token = public.SetTokenUserinfo(u, time.Now().Add(30*24*time.Hour))               //创建一个长效的token
		u.Userstatus = 4                                                                    //给定一个初始的用户权限
		u.Permissions = "user"                                                              //给定默认用户身份为user
		S := time.Now().String()
		u.Signtime = S[0:19]
		R.uid = u.Uid
		R.state = u.Userstatus
		R.token = u.Token
		if !public.VerifyEmailFormat(u.Email) {
			R.msg = 3
		} else {
			R.msg = public.UserinfoAdd(u)
		}

		if R.msg > 1 {
			c.JSON(200, gin.H{
				"code":  200,
				"msg":   R.msg,
				"state": R.state,
			})
		} else {
			//AntiModification := public.MD5(public.MD5(times+u.Uid+u.Token+u.Email) + u.Email + u.Uid + u.Email)
			AntiModification := public.SetTokenUserinfo(u, time.Now().Add(time.Minute*15))
			add := "?uid=" + u.Uid + "&token=" + u.Token + "&email=" + u.Email + "&antimodification=" + AntiModification
			public.SignUp(add)
			c.JSON(200, gin.H{
				"code":  200,
				"msg":   R.msg,
				"state": R.state,
				"token": R.token,
				"uid":   R.uid,
			})
		}
	}
}
