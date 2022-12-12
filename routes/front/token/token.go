package token

import (
	"fmt"
	"gindemo01/common"
	"gindemo01/public/redis"
	"gindemo01/struct/sql_del_struct"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type claimadmin struct {
	UserId     string
	Userstatus string
	Permission string
	jwt.StandardClaims
}
type Claimadmins struct {
	UserId     string
	Userstatus string
	Permission string
}

func SetTokenUserinfo(U sql_del_struct.Userinfo, expireTime time.Duration) string {
	claims := &claimadmin{
		UserId:     U.Uid,
		Userstatus: strconv.Itoa(U.Userstatus),
		Permission: U.Permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "101,43,6,14", // 签名颁发者
			Subject:   "user token",  //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtkey := []byte(common.Jwtinfo.Key)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	// str = tokenString
	err = redis.Set(U.Uid, tokenString, expireTime, 0)
	if err != nil {
		return "reads-error"
	}
	return tokenString
}
func GetTokenUser(tokenString string) Claimadmins {
	var c Claimadmins
	token, c2, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return Claimadmins{
			UserId: "error",
		}
	} else {
		c.UserId = c2.UserId
		c.Userstatus = c2.Userstatus
		c.Permission = c2.Permission
		return c
	}
}
func parseToken(tokenString string) (*jwt.Token, *claimadmin, error) {
	jwtkey := []byte(common.Jwtinfo.Key)
	Claims := &claimadmin{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
func Getting(c *gin.Context) {
	authorizations := c.GetHeader("Authorization")
	cla := GetTokenUser(authorizations)
	if cla.UserId == "error" {
		c.JSON(200, gin.H{
			"code":  200,
			"error": "token",
		})
		c.Abort()
	} else {
		redisToken, err := redis.Get(cla.UserId, 0)
		if err != nil {
			c.JSON(200, gin.H{
				"err": "登录过期",
			})
			c.Abort()
		} else {
			if redisToken == authorizations {
				c.Set("cla", cla)
				err = redis.Set(cla.UserId, redisToken, time.Hour, 0)
				if err != nil {
					c.JSON(200, gin.H{
						"err": "redis",
					})
					c.Abort()
				} else {
					c.Next()
				}
			} else {
				c.JSON(200, gin.H{
					"err": "登录过期",
				})
				c.Abort()
			}
		}
	}
}
