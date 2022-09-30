package sql_operate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// sqlUserId 数据库账号密码
const sqlUserId = "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

func UserInfoFind(u Userinfo) []Userinfo {
	db, err := gorm.Open("mysql", sqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&Userinfo{})
	var user []Userinfo
	db.Debug().Where("uid = ?", u.Uid).First(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
} //查询
func AdminInfoFind(u Admininfo) []Admininfo {
	db, err := gorm.Open("mysql", sqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)
	db.AutoMigrate(&Admininfo{})
	var user []Admininfo
	db.Debug().Where("uid = ?", u.Uid).First(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
}
func UserTokenFind(u Usertoken) []Usertoken {
	db, err := gorm.Open("mysql", sqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)
	u.Updatetime = time.Now().Unix()
	u.Expirationtime = u.Updatetime + 201
	db.AutoMigrate(&Usertoken{})
	var user []Usertoken
	db.Debug().Where("uid = ?", u.Uid).First(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
}
func UserTokenFindTime(u Usertoken) bool {
	U := UserTokenFind(u)
	fmt.Println(U[0].Expirationtime)
	timeS := time.Now().Unix()
	fmt.Println(timeS)
	if timeS > U[0].Expirationtime {
		return false
	}
	return true
} //查询token是否过期
func UserEmailTokenFind(u Useremailtoken) []Useremailtoken {
	db, err := gorm.Open("mysql", sqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&Useremailtoken{})
	var user []Useremailtoken
	db.Debug().Where("uid = ?", u.Uid).First(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
}
