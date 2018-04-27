package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/walle-go/walle-go/controller"
)

func getHandle() *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/user/:id", controller.UserList)
	mux.POST("/user", controller.UserCreate)
	return mux
}
