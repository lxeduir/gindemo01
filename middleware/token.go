package middleware

import (
	"gindemo01/public"
	"gindemo01/routes/front/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenGet(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization") //从header中取Authorization这个key，再与刚刚的token进行比对
	// 先判断取到的是否为空，为空则跳出
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token不能为空"})
		ctx.Abort()
		return
	}
	// 再来解析token，解析失败则跳出
	Uid := login.GetToken(tokenString)
	if Uid == "token错误" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token错误"})
		ctx.Abort()
		return
	}
	// 最后成功了
	var u public.Userinfo
	var find public.Finder = &u
	//U = public.UserInfoFind("uid", Uid, public.Method[0])
	U := find.All("uid", Uid).([]public.Userinfo)
	if len(U) > 0 && U[0].Userstatus == 1 {
		return
	} else if U[0].Userstatus != 1 {
		ctx.JSON(http.StatusOK, gin.H{"code": 401, "uid": Uid, "msg": "账号状态异常"})
		ctx.Abort()
		return
	}
	return
} //验证token是否有效
