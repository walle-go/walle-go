package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/jinzhu/gorm"
	"github.com/walle-go/walle-go/migration"
	_ "github.com/go-sql-driver/mysql"

)

func init() {
	Db, err := gorm.Open("mysql", "root:123456@/walle_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	migration.Migration(Db)
}

func main()  {
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
