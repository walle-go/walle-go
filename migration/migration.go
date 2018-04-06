package migration

import (
	"github.com/walle-go/walle-go/model"
)

func Migration() {
	model.Db.AutoMigrate(&model.User{})
}
