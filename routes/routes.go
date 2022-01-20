package routes

import (
	"database/sql"
	"simple-api/routes/handler"
	"simple-api/utility"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)
}

func Routers(db *sql.DB) {
	Router.HandleFunc("/", utility.HandelRequest(db, handler.Halloworld))
	Router.HandleFunc("/signin", utility.HandelRequest(db, handler.SignIn))
	Router.HandleFunc("/products", utility.HandelRequest(db, handler.SelectAllProduct))
	Router.HandleFunc("/new-product", utility.HandelRequest(db, handler.InsertProduct))
	Router.HandleFunc("/new-log", utility.HandelRequest(db, handler.InsertLog))
	Router.HandleFunc("/logs", utility.HandelRequest(db, handler.SelectLog))
	Router.HandleFunc("/update-log", utility.HandelRequest(db, handler.UpdateLog))
	Router.HandleFunc("/log-history", utility.HandelRequest(db, handler.SelectLogHistory))

}
