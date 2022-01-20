package main

import (
	"fmt"
	"log"
	"net/http"
	ctx "simple-api/models/context"
	"simple-api/routes"
)

var context ctx.DbContext

func main() {
	context = ctx.DbContext{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "hadioz",
		Password: "ozmonday",
		DbName:   "inventory",
	}

	db, err := context.ConnectDB()
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	routes.Routers(db)
	fmt.Println("hallo world")
	if err := http.ListenAndServe(":3000", routes.Router); err != nil {
		log.Fatalf("Server can't be running : %s", err)
	}

}
