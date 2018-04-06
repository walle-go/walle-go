package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	LoginName string
	LoginPwd string
	LoginPwdSalt string
}