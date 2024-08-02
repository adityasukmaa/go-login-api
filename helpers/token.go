package helpers

import (
	"fmt"
	"go-login-api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("mysecretkey")

type MyCustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`

	jwt.RegisteredClaims
}

func CreateToken(user *models.MUser) (string, error) {
	claims := MyCustomClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenStr string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// GenerateImageURL generates a URL to access the user's photo with an API key.
func GenerateImageURL(userID int, fileName string) string {
	baseURL := "http://localhost:8080/users/photo"
	apiKey := "mysecretkey" // replace with the method to retrieve the actual API token
	return fmt.Sprintf("%s/%d/%s?apikey=%s", baseURL, userID, fileName, apiKey)
}
