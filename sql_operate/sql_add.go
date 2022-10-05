package sql_operate

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func UserInfoAdd(U Userinfo) int {
	u := UserInfoFindEmail(U.Email)
	if len(u) > 0 {
		return 2
	} else {
		db, err := gorm.Open("mysql", sqlUserId)
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		if err != nil {
			panic(err)
		}
		db.SingularTable(true) //严格匹配数据库名字
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func AdminInfoAdd(U Admininfo) int {
	u := AdminInfoFind(U)
	if len(u) > 0 {
		return 3
	} else {
		db, err := gorm.Open("mysql", sqlUserId)
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		if err != nil {
			panic(err)
		}
		db.SingularTable(true)
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func UserTokenAdd(U Usertoken) int {
	u := UserTokenFind(U)
	if len(u) > 0 {
		return 2
	} else {
		db, err := gorm.Open("mysql", sqlUserId)
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {
			}
		}(db)
		if err != nil {
			panic(err)
		}
		db.SingularTable(true)
		U.Updatetime = time.Now().Unix()
		U.Expirationtime = U.Updatetime + 201
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func UserEmailTokenAdd(U Useremailtoken) int {
	u := UserEmailTokenFind(U)
	if len(u) > 0 {
		return 2
	} else {
		db, err := gorm.Open("mysql", sqlUserId)
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		if err != nil {
			panic(err)
		}
		db.SingularTable(true) //严格匹配数据库名字
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
