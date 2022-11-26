package public

import (
	"fmt"
	"gindemo01/common"
	"gindemo01/struct/sql_del_struct"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserinfoFind(field string, content string) []sql_del_struct.Userinfo {

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
	db.Debug().Where(field+" = ?", content).Find(&user)
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
func AdmininfoFind(field string, content string) []sql_del_struct.Admininfo {
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
	db.AutoMigrate(&sql_del_struct.Admininfo{})
	var user []sql_del_struct.Admininfo
	db.Debug().Where(field+" = ?", content).Find(&user)
	return user
}
func AdmininfoFirst(field string, content string) sql_del_struct.Admininfo {
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
	db.AutoMigrate(&sql_del_struct.Admininfo{})
	var user []sql_del_struct.Admininfo
	fmt.Println("???")
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
func UserRedisFind(field string, content string) []sql_del_struct.UserRedis {
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
	db.AutoMigrate(&sql_del_struct.UserRedis{})
	var user []sql_del_struct.UserRedis
	db.Debug().Where(field+" = ?", content).Find(&user)
	return user
}
func AdminPathFind() []sql_del_struct.AdminPath {
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
	db.AutoMigrate(&sql_del_struct.AdminPath{})
	var user []sql_del_struct.AdminPath
	db.Debug().Where("rout_id LIKE ?", "%").Find(&user)
	return user
}
func AdminRoleFind() []sql_del_struct.AdminRole {
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
	db.AutoMigrate(&sql_del_struct.AdminRole{})
	var user []sql_del_struct.AdminRole
	db.Debug().Where("role_id LIKE ?", "%").Find(&user)
	return user
}
