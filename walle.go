package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func init() {
	
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
