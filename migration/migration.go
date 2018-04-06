package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/walle-go/walle-go/model"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}