package main

import (
	"net/http"
	"flag"
	"github.com/walle-go/walle-go/migration"
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
	server := &http.Server{
		Addr: "0.0.0.0:9090",
		Handler: getHandle(),
	}
	server.ListenAndServe()
}

func initDatabase()  {
	log.Println("Start database migration.")
	migration.Migration()
	log.Println("Migration end.")
}

func test() {
	log.Println("This is test.")
}
