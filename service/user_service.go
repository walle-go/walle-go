package service

import (
	"github.com/walle-go/walle-go/model"
	"github.com/walle-go/walle-go/repository"
)

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

func (this UserService) Get(id uint) model.User {
	return repository.NewUserRepo().Get(id)
}

func (this UserService) Create() model.User {
	return repository.NewUserRepo().Create()
}