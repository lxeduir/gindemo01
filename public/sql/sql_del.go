package sql

import (
	"gindemo01/struct/sql_del_struct"
	"gindemo01/struct/sql_struct"
)

func DelUserinfo(U sql_struct.Userinfo) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.Userinfo{})
	db.First(&user)
	db.Debug().Where("Uid = ?", U.Uid).Delete(&sql_struct.Userinfo{})
	return 1
}
func DelAdmininfo(uid string) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.Admininfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.Admininfo{})
	db.First(&user)
	db.Debug().Where("Uid = ?", uid).Delete(&sql_struct.Admininfo{})
	return 1
}
func DelUserRedis(U sql_del_struct.UserRedis) int {
	db := GetDB()
	db.SingularTable(true)
	var user sql_del_struct.UserRedis
	db.AutoMigrate(&sql_del_struct.UserRedis{})
	db.First(&user)
	db.Debug().Where("Uid = ? AND Data = ?", U.Uid, U.Data).Delete(&sql_del_struct.UserRedis{})
	return 1
}
func DelAdminRole(U sql_del_struct.AdminRole) int {
	db := GetDB()
	db.SingularTable(true)
	var user sql_del_struct.AdminRole
	db.AutoMigrate(&sql_del_struct.AdminRole{})
	db.First(&user)
	db.Debug().Where("role_id = ?", U.RoleId).Delete(&sql_del_struct.AdminRole{})
	return 1
}
func DelAffaris(U sql_del_struct.Affairs) int {
	db := GetDB()
	db.SingularTable(true)
	var user sql_del_struct.AdminRole
	db.AutoMigrate(&sql_del_struct.AdminRole{})
	db.First(&user)
	db.Debug().Where("affairs_id = ?", U.AffairsId).Delete(&sql_del_struct.AdminRole{})
	return 1
}
