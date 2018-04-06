package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@/walle_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
