package main

import (
	"fmt"
	"go-login-api/configs"
	"go-login-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// app.Get( path: "/swagger/*", swagger.HandlerDefault)
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UsersRoutes(router)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
