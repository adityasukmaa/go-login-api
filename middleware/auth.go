package middleware

import (
	"context"
	"go-login-api/helpers"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			helpers.Response(w, 401, "unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(accessToken)
		if err != nil {
			helpers.Response(w, 401, "unauthorized", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
