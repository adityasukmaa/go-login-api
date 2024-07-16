package helpers

import (
	"fmt"
	"go-login-api/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("mysecretkey")

type MyCustomClaims struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	SessionLogin string `json:"session_login"`
	Gender       string `json:"gender"`
	Photo        string `json:"photo"`
	BirthPlace   string `json:"birth_place"`
	EmployeeID   string `json:"employee_id"`
	EmployeeType string `json:"employee_type"`
	BirthDate    string `json:"birth_date"`

	jwt.RegisteredClaims
}

func CreateToken(user *models.MUser) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Username,
		user.Name,
		user.Email,
		user.PhoneNumber,
		user.SessionLogin,
		user.Gender,
		user.Photo,
		user.BirthPlace,
		user.EmployeeID,
		user.EmployeeType,
		user.BirthDate,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized token")
	}

	return claims, nil
}
