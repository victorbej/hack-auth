package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/victorbej/hack-auth/internal/domain/app"
	"github.com/victorbej/hack-auth/internal/domain/controllers"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/table/new", controllers.CreateUserTable).Methods("POST")
	router.HandleFunc("/api/me/tables", controllers.GetUserTablesFor).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
