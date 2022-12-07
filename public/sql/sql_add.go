package sql

import (
	"gindemo01/struct/sql_struct"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserinfoAdd(U sql_struct.Userinfo) int {
	u := UserinfoFind("email = ?", U.Email)
	if len(u) > 0 {
		return 3
	} else {
		db := GetDB()
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
		db := GetDB()
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
		db := GetDB()
		db.SingularTable(true)
		//fmt.Println(db.NewRecord(&U))
		db.Create(&U) //将上边定义的用户数据写入到数据库user表中
		//fmt.Println(db.NewRecord(&U))
	}
	return 1
}
func UserRedisAdd(U sql_struct.UserRedis) int {
	db := GetDB()
	db.SingularTable(true)
	//fmt.Println(db.NewRecord(&U))
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	//fmt.Println(db.NewRecord(&U))
	return 1
}
func AdminRoleAdd(U sql_struct.AdminRole) int {
	db := GetDB()
	db.SingularTable(true)
	//fmt.Println(db.NewRecord(&U))
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	//fmt.Println(db.NewRecord(&U))
	return 1
}
func AdminPermissionAdd(U sql_struct.AdminPermission) int {
	db := GetDB()
	db.SingularTable(true)
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	return 1
}
func AffairsAdd(U sql_struct.Affairs) int {
	db := GetDB()
	db.SingularTable(true)
	db.Create(&U) //将上边定义的用户数据写入到数据库user表中
	return 1
}
