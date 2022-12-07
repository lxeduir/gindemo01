package login

import (
	"fmt"
	"gindemo01/public/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtkey = []byte("api.edulx.xyz")
var str string

type claims struct {
	UserId string
	jwt.StandardClaims
}

func Setting(ctx *gin.Context) {
	uid, err1 := ctx.GetQuery("uid")
	//U := public.UserInfoFind("uid", uid, public.Method[0])
	U := sql.UserinfoFind("uid = ?", uid)
	if len(U) == 0 {
		ctx.JSON(201, gin.H{
			"msg": "uid不存在",
		})
		return
	}
	if !err1 {
		ctx.JSON(201, gin.H{
			"msg": "缺少必需参数",
		})
		return
	}
	expireTime := time.Now().Add(10 * time.Minute)
	claims := &claims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "101.43.6.14", // 签名颁发者
			Subject:   "user token",  //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	// str = tokenString
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "token": tokenString})
}
func Getting(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization") //从header中取Authorization这个key，再与刚刚的token进行比对
	// 先判断取到的是否为空，为空则跳出
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "token为空"})
		ctx.Abort()
		return
	}
	// 再来解析token，解析失败则跳出
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "权限不足"})
		ctx.Abort()
		return
	}
	// 最后成功了
	ctx.JSON(http.StatusOK, gin.H{"id": claims.UserId, "msg": "认证通过"})

}
func ParseToken(tokenString string) (*jwt.Token, *claims, error) {
	Claims := &claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
func GetToken(tokenString string) string {
	if tokenString == "" {
		return "token不能为空"
	}
	// 再来解析token，解析失败则跳出
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return "token错误"
	}
	// 最后成功了
	return claims.UserId
}
func SetToken(Uid string) string {
	U := sql.UserinfoFind("uid = ?", Uid)
	if len(U) == 0 {
		return "用户不存在"
	}
	expireTime := time.Now().Add(10 * time.Minute)
	claims := &claims{
		UserId: Uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "101,43,6,14", // 签名颁发者
			Subject:   "user token",  //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	// str = tokenString
	return tokenString
}
