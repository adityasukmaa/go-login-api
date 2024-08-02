package middleware

import (
	"context"
	"go-login-api/helpers"
	"log"
	"net/http"
)

// ApiKeyAuth middleware validates the API key present in the URL query parameters.
func ApiKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.URL.Query().Get("apikey")
		if apiKey == "" {
			log.Println("Missing API key")
			helpers.Response(w, http.StatusUnauthorized, "Unauthorized: Missing API Key", nil)
			return
		}

		// Validate API Key
		user, err := helpers.ValidateToken(apiKey)
		if err != nil {
			log.Printf("Invalid API key: %v", err)
			helpers.Response(w, http.StatusUnauthorized, "Unauthorized: Invalid API Key", nil)
			return
		}

		// Log the user info for debugging
		log.Printf("Valid API key for user ID %d", user.ID)

		// Store user info in context
		ctx := context.WithValue(r.Context(), "userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

