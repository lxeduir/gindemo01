package public

import "github.com/jinzhu/gorm"

type Finder interface {
	First(field string, content string) interface{}
	All(field string, content string) interface{}
}

func (*Userinfo) First(field string, content string) interface{} {
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&Userinfo{})
	var user []Userinfo
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
func (*Userinfo) All(field string, content string) interface{} {
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&Userinfo{})
	var user []Userinfo
	db.Debug().Where(field+" = ?", content).Find(&user)
	return user
}
func (*Admininfo) First(field string, content string) interface{} {
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&Admininfo{})
	var user []Admininfo
	db.Debug().Where(field+" = ?", content).First(&user)
	return user[0]
}
func (*Admininfo) All(field string, content string) interface{} {
	db, err := gorm.Open("mysql", SqlUserId)
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.SingularTable(true) //严格匹配数据库名字
	db.AutoMigrate(&Admininfo{})
	var user []Admininfo
	db.Debug().Where(field+" = ?", content).Find(&user)
	return user
}
