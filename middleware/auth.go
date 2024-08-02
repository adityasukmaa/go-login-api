package middleware

import (
	"context"
	"go-login-api/helpers"
	"log"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			log.Println("Authorization header is missing")
			helpers.Response(w, 401, "unauthorized", nil)
			return
		}

		tokenParts := strings.Split(accessToken, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Println("Authorization header format must be Bearer {token}")
			helpers.Response(w, 401, "unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(tokenParts[1])
		if err != nil {
			log.Printf("Token validation failed: %v", err)
			helpers.Response(w, 401, "unauthorized", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
