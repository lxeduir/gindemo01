package public

import (
	"gindemo01/common"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open("mysql", common.MysqlInfo.Id)
	if err != nil {
		return err
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.SingularTable(true)
	return nil

}
func CloseDB() {
	err := DB.Close()
	if err != nil {
		return
	}
}
func GetDB() *gorm.DB {
	return DB
}
