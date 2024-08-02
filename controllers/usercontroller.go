package controllers

import (
	"encoding/json"
	"fmt"
	"go-login-api/configs"
	"go-login-api/helpers"
	"go-login-api/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Get Profile
func Profile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	// Log the request
	log.Printf("User ID %d with username %s is accessing profile", user.ID, user.Username)

	var userProfile models.MUser
	if err := configs.DB.First(&userProfile, user.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.Response(w, 404, "User not found", nil)
		} else {
			helpers.Response(w, 500, "Error retrieving user", nil)
		}
		return
	}

	userResponse := &models.MyProfile{
		ID:          userProfile.ID,
		Name:        userProfile.Name,
		Username:    userProfile.Username,
		Email:       userProfile.Email,
		PhoneNumber: userProfile.PhoneNumber,
		Gender:      userProfile.Gender,
		BirthPlace:  userProfile.BirthPlace,
		BirthDate:   userProfile.BirthDate,
		PhotoURL:    userProfile.PhotoURL,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}

// Update Profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	var updatedData models.MUser

	// Log the incoming payload
	log.Println("Received payload for update profile")

	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		helpers.Response(w, 400, "Invalid request payload", nil)
		return
	}

	// Log the decoded payload
	log.Printf("Decoded payload: %+v", updatedData)

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

	userResponse := &models.MyProfile{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		BirthPlace:  user.BirthPlace,
		BirthDate:   user.BirthDate,
	}

	helpers.Response(w, 200, "User profile", userResponse)
}

// Upload Foto
func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	// Parse form to retrieve file
	err := r.ParseMultipartForm(10 << 20) // Max 10 MB
	if err != nil {
		helpers.Response(w, 400, "Unable to parse form", nil)
		return
	}

	file, handler, err := r.FormFile("photo")
	if err != nil {
		helpers.Response(w, 400, "Unable to retrieve file", nil)
		return
	}
	defer file.Close()

	// Create a unique file name
	fileName := fmt.Sprintf("%d_%s", user.ID, filepath.Base(handler.Filename))
	filePath := filepath.Join("public/uploads", fileName)

	// Buat direktori jika belum ada
	if _, err := os.Stat("public/uploads"); os.IsNotExist(err) {
		os.Mkdir("public/uploads", os.ModePerm)
	}

	// Create the file on server
	dst, err := os.Create(filePath)
	if err != nil {
		helpers.Response(w, 500, "Unable to create file", nil)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := dst.ReadFrom(file); err != nil {
		helpers.Response(w, 500, "Unable to save file", nil)
		return
	}

	// Update the user's photo in the database
	if err := configs.DB.Model(&models.MUser{}).Where("id = ?", user.ID).Update("photo", fileName).Error; err != nil {
		helpers.Response(w, 500, "Failed to update user photo", nil)
		return
	}

	helpers.Response(w, 200, "Photo uploaded successfully", nil)
}

// Get Photo
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	vars := mux.Vars(r)
	photoName := vars["photo"]

	// Check if the photo belongs to the user
	var userProfile models.MUser
	if err := configs.DB.First(&userProfile, user.ID).Error; err != nil {
		helpers.Response(w, 500, "Error retrieving user", nil)
		return
	}

	if userProfile.Photo != photoName {
		helpers.Response(w, 403, "Access denied", nil)
		return
	}

	photoPath := filepath.Join("public/uploads", photoName)

	// Check if file exists
	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		helpers.Response(w, 404, "Photo not found", nil)
		return
	}

	// Serve the file
	http.ServeFile(w, r, photoPath)
}

// Delete Photo
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	// Get user profile from DB
	var userProfile models.MUser
	if err := configs.DB.First(&userProfile, user.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.Response(w, 404, "User not found", nil)
		} else {
			helpers.Response(w, 500, "Error retrieving user", nil)
		}
		return
	}

	if userProfile.Photo == "" {
		helpers.Response(w, 400, "No photo to delete", nil)
		return
	}

	// Construct photo path
	photoPath := filepath.Join("public/uploads", userProfile.Photo)

	// Delete the photo file from server
	if err := os.Remove(photoPath); err != nil {
		helpers.Response(w, 500, "Failed to delete photo", nil)
		return
	}

	// Update the user's photo in the database to be empty
	if err := configs.DB.Model(&models.MUser{}).Where("id = ?", user.ID).Update("photo", "").Error; err != nil {
		helpers.Response(w, 500, "Failed to update user photo", nil)
		return
	}

	helpers.Response(w, 200, "Photo deleted successfully", nil)
}
