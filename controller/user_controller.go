package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/walle-go/walle-go/service"
	"strconv"
)

func UserList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		panic(err)
	}
	user := service.NewUserService().Get(uint(id))
	fmt.Fprintf(w, "hello, %s!", user.Name)
}

func UserCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	service.NewUserService().Create()
	fmt.Fprintf(w, "success")
}
