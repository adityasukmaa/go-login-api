package routes

import (
	"go-login-api/controllers"
	"go-login-api/middleware"

	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/profile", controllers.Profile).Methods("GET")
	router.HandleFunc("/profile", controllers.UpdateProfile).Methods("PUT")
	// router.HandleFunc("/profile", controllers.DeleteProfile).Methods("DELETE")
	// router.HandleFunc("/profile/{id}", controllers.GetProfileByID).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

}
