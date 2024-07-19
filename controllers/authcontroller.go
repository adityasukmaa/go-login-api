package controllers

import (
	"encoding/json"
	"go-login-api/configs"
	"go-login-api/helpers"
	"go-login-api/models"
	"log"
	"net/http"
	"net/http/httputil"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	dump, _ := httputil.DumpRequest(r, true)
	log.Printf("req: %q\n", dump)

	// Check if the username or email already exists
	var existingUser models.MUser
	if err := configs.DB.Where("username = ? OR email = ?", register.Username, register.Email).First(&existingUser).Error; err == nil {
		helpers.Response(w, 400, "Username or Email already exists", nil)
		return
	}

	// if register.Password != register.PasswordConfirm {
	// 	helpers.Response(w, 400, "Password not match", nil)
	// 	return
	// }

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.MUser{
		Username: register.Username,
		Password: passwordHash,
		Name:     register.Name,
		Email:    register.Email,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Register Succesfully", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.MUser
	if err := configs.DB.First(&user, "username = ?", login.Username).Error; err != nil {
		helpers.Response(w, 404, "Wrong Username or Password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong Username or Password", nil)
		return
	}

	dump, _ := httputil.DumpRequest(r, true)
	log.Printf("req: %q\n", dump)

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Successfully Login", token)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		helpers.Response(w, 400, "No token provided", nil)
		return
	}

	// Optional: Add the token to a blacklist if you are managing blacklisted tokens
	// blacklistToken(token)

	helpers.Response(w, 200, "Successfully logged out", nil)
}
