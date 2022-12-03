package public

import (
	"gindemo01/common"
	"gindemo01/struct/sql_struct"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserinfoAdd(U sql_struct.Userinfo) int {
	u := UserinfoFind("email = ?", U.Email)
	if len(u) > 0 {
		return 3
	} else {
		db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
func AdmininfoAdd(U sql_struct.Admininfo) int {
	u := AdmininfoFind("email = ?", U.Email)
	if len(u) > 0 {
		return 3
	} else {
		db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
func UserImgAdd(U sql_struct.UserImg) int {
	u := UserinfoFind("uid = ?", U.Uid)
	if len(u) == 0 {
		return 3
	} else {
		db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
func UserRedisAdd(U sql_struct.UserRedis) int {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
	return 1
}
func AdminRoleAdd(U sql_struct.AdminRole) int {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
	return 1
}
func AdminPermissionAdd(U sql_struct.AdminPermission) int {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	return 1
}
