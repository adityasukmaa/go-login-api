package main

import (
	"fmt"
	"go-login-api/configs"
	"go-login-api/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// app.Get( path: "/swagger/*", swagger.HandlerDefault)
	configs.ConnectDB()

	allowedOrigins := []string{"*"}
	allowedMethods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{"Content-Type", "Authorization"}

	cors := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedMethods(allowedMethods),
	)

	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(cors)

	router := r.PathPrefix("/api").Subrouter().StrictSlash(true)
	router2 := r.PathPrefix("/apiv2").Subrouter().StrictSlash(true)

	// Menambahkan middleware CORS ke dalam sub-router `/api`
	// corsHandler := cors.New(cors.Options{
	// 	AllowedOrigins:   allowedOrigins, // Mengizinkan semua origin
	// 	AllowedMethods:   allowedMethods,
	// 	AllowedHeaders:   allowedHeaders,
	// 	AllowCredentials: true,
	// })
	// router.Use(corsHandler.Handler)

	// Menambahkan rute untuk ping
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "xaviera cantik"}`))
	}).Methods("GET")

	routes.AuthRoutes(router)
	routes.UsersRoutes(router)
	routes.AuthRoutes(router2)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe("0.0.0.0:8080", cors(r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
