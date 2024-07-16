package controllers

import (
	"encoding/json"
	"go-login-api/configs"
	"go-login-api/helpers"
	"go-login-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Get Profile
func Profile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	userResponse := &models.MyProfile{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		SessionLogin: user.SessionLogin,
		Gender:       user.Gender,
		Photo:        user.Photo,
		BirthPlace:   user.BirthPlace,
		EmployeeID:   user.EmployeeID,
		EmployeeType: user.EmployeeType,
		BirthDate:    user.BirthDate,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}

// Update Profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	var updatedData models.MUser

	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		helpers.Response(w, 400, "Invalid request payload", nil)
		return
	}

	if err := configs.DB.Model(&models.MUser{}).Where("id = ?", user.ID).Updates(updatedData).Error; err != nil {
		helpers.Response(w, 500, "Failed to update profile", nil)
		return
	}

	helpers.Response(w, 200, "Profile updated successfully", nil)
}

// Delete Profile
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	if err := configs.DB.Where("id = ?", user.ID).Delete(&models.MUser{}).Error; err != nil {
		helpers.Response(w, 500, "Failed to delete profile", nil)
		return
	}

	helpers.Response(w, 200, "Profile deleted successfully", nil)
}

// Get Profile by ID
func GetProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.Response(w, 400, "Invalid user ID", nil)
		return
	}

	var user models.MUser
	if err := configs.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.Response(w, 404, "User not found", nil)
		} else {
			helpers.Response(w, 500, "Error retrieving user", nil)
		}
		return
	}

	helpers.Response(w, 200, "User profile", user)
}
