package main

import (
	"database/sql"

	"./controllers"
	"./drivers"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()

	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/protected", controller.TokenVerifyMiddleware(controller.ProtectedEndpoint())).Methods("GET")

}
