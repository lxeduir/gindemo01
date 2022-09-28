package sql_operate

import (
	"github.com/jinzhu/gorm"
)

type Userinfos struct {
	Uid    string
	Name   string
	Passwd string
	Token  string
} //查找数据时用结构体
type Admininfos struct {
	Uid    string
	Name   string
	Passwd string
	Token  string
	Rtoken string
} //插入数据用结构体
func UserInfoFind(u Userinfo) []Userinfos {
	db, err := gorm.Open("mysql", "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)
	db.AutoMigrate(&Userinfos{})
	var user []Userinfos
	db.Debug().Where("uid = ?", u.Uid).First(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
}
func AdminInfoFind(u Admininfo) []Admininfos {
	db, err := gorm.Open("mysql", "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)
	db.AutoMigrate(&Admininfos{})
	var user []Admininfos
	db.Debug().Where("uid = ?", u.Uid).First(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
}

//func UserInfoFindPass(u Userinfo) []Userinfos {
//	db, err := gorm.Open("mysql", "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		panic(err)
//	}
//	defer func(db *gorm.DB) {
//		err := db.Close()
//		if err != nil {
//		}
//	}(db)
//
//	db.SingularTable(true)
//	db.AutoMigrate(&Userinfos{})
//	var user []Userinfos
//	db.Debug().Where("uid = ?", u.Uid).First(&user)
//	//fmt.Println("查询第一条匹配条件记录：", user)
//	return user
//}
