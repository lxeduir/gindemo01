package public

import (
	"gindemo01/config"
	"github.com/jinzhu/gorm"
	"time"
)

var Method = []string{"first", "contain", "all"}

func UserInfoFind(target string, content string, method string) []config.Userinfo {
	db, err := gorm.Open("mysql", config.SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&config.Userinfo{})
	var user []config.Userinfo
	if method == "first" {
		db.Debug().Where(target+" = ?", content).First(&user) //查询第一条匹配条件记录
	} else if method == "contain" {
		db.Debug().Where(target+" LIKE ?", "%"+content+"%").Find(&user) //查询所有包含contrnt的记录
	} else if method == "all" {
		db.Debug().Where(target+" = ?", content).Find(&user) //查询所有匹配条件的记录
	}
	return user
} //userinfo表通用查询方法
func AdminInfoFind(target string, content string, method string) []config.Admininfo {
	db, err := gorm.Open("mysql", config.SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)
	db.AutoMigrate(&config.Admininfo{})
	var user []config.Admininfo
	if method == "first" {
		db.Debug().Where(target+" = ?", content).First(&user) //查询第一条匹配条件记录
	} else if method == "contain" {
		db.Debug().Where(target+" LIKE ?", "%"+content+"%").Find(&user) //查询所有包含contrnt的记录
	} else if method == "all" {
		db.Debug().Where(target+" = ?", content).Find(&user) //查询所有匹配条件的记录
	}
	return user
}
func UserTokenFind(u config.Usertoken) []config.Usertoken {
	db, err := gorm.Open("mysql", config.SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true)
	u.Updatetime = time.Now().Unix()
	u.Expirationtime = u.Updatetime + 600
	db.AutoMigrate(&config.Usertoken{})
	var user []config.Usertoken
	db.Debug().Where("uid = ?", u.Uid).Find(&user)
	return user
}
func UserEmailTokenFind(u config.Useremailtoken) []config.Useremailtoken {
	db, err := gorm.Open("mysql", config.SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&config.Useremailtoken{})
	var user []config.Useremailtoken
	db.Debug().Where("uid = ?", u.Uid).Find(&user)
	//fmt.Println("查询第一条匹配条件记录：", user)
	return user
}
