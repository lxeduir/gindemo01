package public

import (
	"fmt"
	"gindemo01/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func UserInfoReviseStates(U config.Userinfo) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", config.SqlUserId)
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
	var user config.Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&config.Userinfo{})
	db.First(&user)
	fmt.Println("********************************")
	fmt.Println(user)
	fmt.Println("********************************")
	db.Debug().Model(&user).Where("uid = ?", U.Uid).Update("Userstatus", U.Userstatus)
}
func AdminInfoRevise(U config.Admininfo) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", config.SqlUserId)
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
	db.AutoMigrate(&config.Userinfo{})
	var user config.Userinfo
	db.First(&user)
	fmt.Println(user)
	db.Debug().Model(&user).Updates(gin.H{
		"Uid":    U.Uid,
		"Passwd": U.Passwd,
		"Token":  U.Token,
	})
}
func UserTokenRevise(U config.Usertoken) string {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", config.SqlUserId)
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
	db.AutoMigrate(&config.Usertoken{})
	var user config.Usertoken
	db.First(&user)
	db.Debug().Model(&user).Updates(U)
	return U.Token
} //更新token