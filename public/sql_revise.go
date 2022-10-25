package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func UserInfoReviseStates(U Userinfo) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&Userinfo{})
	db.First(&user)
	fmt.Println("********************************")
	fmt.Println(user)
	fmt.Println("********************************")
	db.Debug().Model(&user).Where("uid = ?", U.Uid).Update("Userstatus", U.Userstatus)
}
func AdminInfoRevise(U Admininfo) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)

	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&Userinfo{})
	var user Userinfo
	db.First(&user)
	fmt.Println(user)
	db.Debug().Model(&user).Updates(gin.H{
		"Uid":    U.Uid,
		"Passwd": U.Passwd,
		"Token":  U.Token,
	})
}
func UserTokenRevise(U Usertoken) string {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	U.Token = MD5(U.Token + strconv.FormatInt(time.Now().Unix(), 10))
	U.Updatetime = time.Now().Unix()
	U.Expirationtime = U.Updatetime + 201
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&Usertoken{})
	var user Usertoken
	db.First(&user)
	db.Debug().Model(&user).Updates(U)
	return U.Token
} //更新token
