package routes

import (
	"go-login-api/controllers"
	"go-login-api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/profile", controllers.Profile).Methods("GET")
	router.HandleFunc("/profile", controllers.UpdateProfile).Methods("PUT")
	// router.HandleFunc("/profile", controllers.DeleteProfile).Methods("DELETE")
	router.HandleFunc("/profile/{id}", controllers.GetProfileByID).Methods("GET")
	router.HandleFunc("/profile/photo", controllers.UploadPhoto).Methods("POST")
	router.HandleFunc("/profile/photo/{photo}", controllers.GetPhoto).Methods("GET")
	router.HandleFunc("/profile/photo", controllers.DeletePhoto).Methods("DELETE")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Serve static files from /public/uploads directory with API key auth
	fs := http.FileServer(http.Dir("./public/uploads"))
	router.PathPrefix("/photo").Handler(http.StripPrefix("/photo", middleware.ApiKeyAuth(fs)))
}
