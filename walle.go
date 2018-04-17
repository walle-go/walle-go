package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/walle-go/walle-go/controller"
	"flag"
	"github.com/walle-go/walle-go/migration"
	"github.com/walle-go/walle-go/repository"
	"log"
)

func main()  {
	var isHelp,isInit,isTest bool
	flag.BoolVar(&isHelp, "h", false, "this help")
	flag.BoolVar(&isInit, "i", false, "init database")
	flag.BoolVar(&isTest, "t", false, "test sql")
	flag.Parse()

	if isHelp {
		flag.Usage()
	} else if isInit {
		initDatabase()
	} else if isTest {
		test()
	} else {
		serve()
	}
}

func serve() {
	log.Println("Start serve...")
	mux := httprouter.New()
	mux.GET("/user/:id", controller.UserList)
	server := &http.Server{
		Addr: "0.0.0.0:9090",
		Handler: mux,
	}
	server.ListenAndServe()
}

func initDatabase()  {
	log.Println("Start database migration.")
	migration.Migration()
	log.Println("Migration end.")
}

func test() {
	log.Println("Start insert into test data.")
	var userRepo repository.UserRepo
	userRepo.InsertOne()
	log.Println("Insert into end.")
}
