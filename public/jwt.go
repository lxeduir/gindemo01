package public

import (
	"fmt"
	"gindemo01/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtkey = []byte("api.edulx.xyz")
var str string

type claims struct {
	UserId      string
	Permissions string
	Userstatus  string
	jwt.StandardClaims
}

func parseToken(tokenString string) (*jwt.Token, *claims, error) {
	Claims := &claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
func GetToken(tokenString string) gin.H {
	if tokenString == "" {
		return gin.H{
			"msg": "token不能为空",
		}
	}
	// 再来解析token，解析失败则跳出
	token, claims, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return gin.H{
			"msg": "token错误",
		}
	}
	// 最后成功了
	U := UserInfoFind("uid", claims.UserId, Method[0])
	if len(U) == 0 {
		return gin.H{
			"msg": "用户不存在",
		}
	}
	return gin.H{
		"mas": 1,
	}
}

func SetToken(U config.Userinfo, expireTime time.Time) string {
	//expireTime := time.Now().Add(24 * time.Hour)
	claims := &claims{
		UserId:      U.Uid,
		Permissions: U.Permissions,
		Userstatus:  string(U.Userstatus),
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
