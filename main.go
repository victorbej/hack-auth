package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"Access-Control-Allow-Origin", "*"},
		AllowedMethods:   []string{"Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
