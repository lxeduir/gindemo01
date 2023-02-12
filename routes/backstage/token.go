package backstage

import (
	"fmt"
	"gindemo01/common"
	"gindemo01/public"
	"gindemo01/public/redis"
	"gindemo01/struct/sql_del_struct"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var jwtkey = []byte("")

type claimadmin struct {
	UserId     string
	RoleId     string
	Userstatus string
	Mps        map[string]map[int]int
	jwt.StandardClaims
}
type Claimadmins struct {
	UserId     string
	RoleId     string
	Userstatus string
	Mps        map[string]map[int]int
}

func SetTokenAdmininfo(U sql_del_struct.Admininfo, expireTime time.Duration) string {
	mp := public.Authmap(strconv.Itoa(U.RoleId))
	claims := &claimadmin{
		UserId:     U.Uid,
		RoleId:     strconv.Itoa(U.RoleId),
		Userstatus: strconv.Itoa(U.State),
		Mps:        mp,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "101,43,6,14", // 签名颁发者
			Subject:   "admin token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtkey = []byte(common.Jwtinfo.Key)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	// str = tokenString
	err = redis.Set(U.Uid, tokenString, expireTime, 1)
	if err != nil {
		return "reads-error"
	}
	return tokenString
}
func GetTokenAdmin(tokenString string) Claimadmins {
	var c Claimadmins
	token, c2, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return Claimadmins{
			UserId: "error",
		}
	} else {
		c.UserId = c2.UserId
		c.Userstatus = c2.Userstatus
		c.RoleId = c2.RoleId
		c.Mps = c2.Mps
		return c
	}
}
func parseToken(tokenString string) (*jwt.Token, *claimadmin, error) {
	jwtkey = []byte(common.Jwtinfo.Key)
	Claims := &claimadmin{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
func getting(c *gin.Context) {
	authorizations := c.GetHeader("Authorization")
	cla := GetTokenAdmin(authorizations)
	if cla.UserId == "error" {
		c.JSON(200, gin.H{
			"err": "token",
		})
		c.Abort()
	} else {
		redisToken, err := redis.Get(cla.UserId, 1)
		if err != nil {
			c.JSON(200, gin.H{
				"err": "登录过期",
			})
			c.Abort()
		} else {
			if redisToken == authorizations {
				c.Set("cla", cla)
				err = redis.Set(cla.UserId, redisToken, time.Hour, 1)
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
