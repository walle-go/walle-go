package repository

import "github.com/walle-go/walle-go/model"

type UserRepo struct {

}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (this UserRepo) Get(id uint) (user model.User) {
	model.Db.Take(&user, id)
	return
}

func (this UserRepo) List() {

}

func (this UserRepo) Create() model.User {
	var user model.User
	user.Name = "liuyibao"
	model.Db.NewRecord(user)
	model.Db.Create(&user)
	return user
}

func (this UserRepo) Update() {

}

func (this UserRepo) Delete() {

}