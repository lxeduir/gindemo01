package sql

import (
	"gindemo01/struct/sql_struct"
	"time"
)

func ReviseUserinfo(U sql_struct.Userinfo) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.Userinfo{})
	db.First(&user)
	db.Debug().Model(&user).Where("uid = ?", U.Uid).Updates(map[string]interface{}{"Username": U.Username, "Permissions": U.Permissions, "Passwd": U.Passwd, "Userstatus": U.Userstatus})
	return 1
}
func ReviseAdmininfo(U sql_struct.Admininfo) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.Admininfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.Admininfo{})
	db.First(&user)
	db.Debug().Model(&user).Where("uid = ?", U.Uid).Updates(map[string]interface{}{"Username": U.Username, "State": U.State, "Passwd": U.Passwd, "Token": U.Token})
	return 1
}
func ReviseAdminrole(U sql_struct.AdminRole) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.AdminRole
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.AdminRole{})
	db.First(&user)
	db.Debug().Model(&user).Where("role_id = ?", U.RoleId).Updates(map[string]interface{}{
		"name":            U.Name,
		"description":     U.Description,
		"update_time":     time.Now().String()[0:19],
		"update_by":       U.UpdateBy,
		"permission_json": U.PermissionJson,
		"orders":          U.Orders,
	})
	return 1
}
func ReviseAdminpermission(U sql_struct.AdminPermission) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.AdminPermission
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.AdminPermission{})
	db.First(&user)
	db.Debug().Model(&user).Where("permission_id = ?", U.PermissionId).Updates(map[string]interface{}{})
	return 1
}
func ReviseAffairs(U sql_struct.Affairs) int {
	db := GetDB()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	var user sql_struct.AdminPermission
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.AdminPermission{})
	db.First(&user)
	db.Debug().Model(&user).Where("affairs_id = ?", U.AffairsId).Updates(map[string]interface{}{
		"affairs_data": U.AffairsData,
		"state":        U.State,
		"update_time":  U.UpdateTime,
		"update_by":    U.UpdateBy,
		"dispose_time": U.DisposeTime,
	})
	return 1
}
