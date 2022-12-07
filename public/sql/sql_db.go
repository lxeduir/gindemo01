package sql

import (
	"gindemo01/common"
	"github.com/jinzhu/gorm"
)

var dB *gorm.DB

func InitDB() (err error) {
	dB, err = gorm.Open("mysql", common.MysqlInfo.Id)
	if err != nil {
		return err
	}
	dB.DB().SetMaxIdleConns(10)
	dB.DB().SetMaxOpenConns(100)
	dB.SingularTable(true)
	return nil
}
func CloseDB() {
	err := dB.Close()
	if err != nil {
		return
	}
}
func GetDB() *gorm.DB {
	return dB

}
