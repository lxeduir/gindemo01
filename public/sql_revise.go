package public

import (
	"gindemo01/common"
	"gindemo01/struct/sql_struct"
	"github.com/jinzhu/gorm"
	"time"
)

func ReviseUserinfo(U sql_struct.Userinfo) int {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
	var user sql_struct.Userinfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.Userinfo{})
	db.First(&user)
	db.Debug().Model(&user).Updates(map[string]interface{}{"Username": U.Username, "Permissions": U.Permissions, "Passwd": U.Passwd, "Userstatus": U.Userstatus})
	return 1
}
func ReviseAdmininfo(U sql_struct.Admininfo) int {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
	var user sql_struct.Admininfo
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.Admininfo{})
	db.First(&user)
	db.Debug().Model(&user).Updates(map[string]interface{}{"Username": U.Username, "State": U.State, "Passwd": U.Passwd, "Token": U.Token})
	return 1
}
func ReviseAdminrole(U sql_struct.AdminRole) int {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
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
	var user sql_struct.AdminPermission
	// 3, 把模型与数据库中的表对应起来
	db.AutoMigrate(&sql_struct.AdminPermission{})
	db.First(&user)
	db.Debug().Model(&user).Where("permission_id = ?", U.PermissionId).Updates(map[string]interface{}{})
	return 1
}
