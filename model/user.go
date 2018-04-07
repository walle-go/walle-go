package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null; default:''"`
	LoginName string `gorm:"not null; default:''"`
	LoginPwd string `gorm:"not null; default:''"`
	LoginPwdSalt string `gorm:"not null; default:''"`
}