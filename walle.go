package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/walle-go/walle-go/model"
	"github.com/walle-go/walle-go/migration"
)

func main()  {
	migration.Migration()
	var user model.User
	user.Name = "liuyibao"
	model.InsertOne(user)
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	server := &http.Server{
		Addr: "0.0.0.0:9090",
		Handler: mux,
	}
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!", p.ByName("name"))
}
