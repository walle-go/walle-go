package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/walle-go/walle-go/service"
	"strconv"
	"encoding/json"
	"github.com/walle-go/walle-go/util"
)

func UserList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		panic(err)
	}
	user := service.NewUserService().Get(uint(id))
	userByte, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set(util.GetJsonHeader())
	w.Write(userByte)
}

func UserCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	service.NewUserService().Create()
	fmt.Fprintf(w, "success")
}
