package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/walle-go/walle-go/migration"
	"github.com/walle-go/walle-go/repository"
	"github.com/walle-go/walle-go/controller"
)

func main()  {
	migration.Migration()
	var userRepo repository.UserRepo
	userRepo.InsertOne()
	mux := httprouter.New()
	mux.GET("/user/:id", controller.UserList)
	server := &http.Server{
		Addr: "0.0.0.0:9090",
		Handler: mux,
	}
	server.ListenAndServe()
}
