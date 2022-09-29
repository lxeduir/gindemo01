package sql_operate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func UserInfoRevise(U Userinfo) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", sqlUserId)
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
		"Name":   U.Name,
		"Passwd": U.Passwd,
	})
}
func AdminInfoRevise(U Admininfo) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", sqlUserId)
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
		"Name":   U.Name,
		"Passwd": U.Passwd,
		"Token":  U.Token,
		"Rtoken": U.Rtoken,
	})
}
func UserTokenRevise(U Usertoken) {
	// 2, 连接数据库
	db, err := gorm.Open("mysql", sqlUserId)
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
	U.Updatetime = time.Now().Unix()
	U.Expirationtime = U.Updatetime + 201
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&Usertoken{})
	var user Usertoken
	db.First(&user)
	fmt.Println(user)
	db.Debug().Model(&user).Updates(U)
} //刷新token过期时间
