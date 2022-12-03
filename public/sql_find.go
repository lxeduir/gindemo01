package public

import (
	"gindemo01/common"
	"gindemo01/struct/sql_del_struct"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserinfoFind(query string, args string) []sql_del_struct.Userinfo {

	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.Userinfo{})
	var user []sql_del_struct.Userinfo
	db.Debug().Where(query, args).Find(&user)
	return user
}
func UserinfoFirst(field string, content string) sql_del_struct.Userinfo {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.Userinfo{})
	var user []sql_del_struct.Userinfo
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
func AdmininfoFind(query string, args string) []sql_del_struct.Admininfo {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.Admininfo{})
	var user []sql_del_struct.Admininfo
	db.Debug().Where(query, args).Find(&user)
	return user
}
func AdmininfoFirst(field string, content string) sql_del_struct.Admininfo {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.Admininfo{})
	var user []sql_del_struct.Admininfo
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
func UserRedisFind(field string, content string) []sql_del_struct.UserRedis {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.UserRedis{})
	var user []sql_del_struct.UserRedis
	db.Debug().Where(field+" = ?", content).Find(&user)
	return user
}
func AdminRoutFind(query string, args string) []sql_del_struct.AdminRout {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.AdminRout{})
	var user []sql_del_struct.AdminRout
	db.Debug().Where(query, args).Find(&user)
	return user
}
func AdminRoleFind(query string, args string) []sql_del_struct.AdminRole {
	db := GetDB()
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.AdminRole{})
	var user []sql_del_struct.AdminRole
	db.Debug().Where(query, args).Find(&user)
	return user
}
func AdminPermissionFind(query string, args string) []sql_del_struct.AdminPermission {
	db, err := gorm.Open("mysql", common.MysqlInfo.Id)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&sql_del_struct.AdminPermission{})
	var user []sql_del_struct.AdminPermission
	db.Debug().Where(query, args).Find(&user)
	return user
}
