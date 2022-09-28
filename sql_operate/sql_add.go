package sql_operate

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Userinfo struct {
	Uid    string
	Name   string
	Passwd string
	Token  string
} //插入数据用结构体
type Admininfo struct {
	Uid    string
	Name   string
	Passwd string
	Token  string
	Rtoken string
} //插入数据用结构体

//	func Main() {
//		db, err := gorm.Open("mysql", "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
//		defer func(db *gorm.DB) {
//			err := db.Close()
//			if err != nil {
//
//			}
//		}(db)
//		if err != nil {
//			panic(err)
//		}
//	}
func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
func UserInfoAdd(U Userinfo) int {
	u := UserInfoFind(U)
	if len(u) > 0 {
		return 3
	} else {
		db, err := gorm.Open("mysql", "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		if err != nil {
			panic(err)
		}
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
} //userinfo表单
func AdminInfoAdd(U Admininfo) int {
	u := AdminInfoFind(U)
	if len(u) > 0 {
		return 3
	} else {
		db, err := gorm.Open("mysql", "users:liuxun@(101.43.6.142:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		if err != nil {
			panic(err)
		}
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
