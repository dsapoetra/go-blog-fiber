package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-blog-fiber/app/model/http"
	"go-blog-fiber/app/model/repo"
	"os"
	"time"
)

func SignedToken(author repo.Author) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(3))),
		Issuer:    "test",
		ID:        author.ID.String(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*http.PayloadJWT, error) {
	token, err := jwt.ParseWithClaims(tokenString, &http.PayloadJWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("err here")
	}

	// type-assert `Claims` into a variable of the appropriate type
	myClaims := token.Claims.(*http.PayloadJWT)

	return myClaims, nil

}
