package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-blog-fiber/app/model/repo"
	"log"
	"os"
	"time"
)

func SignedToken(author repo.Author) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
		ID:        author.ID.String(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println("HERE 3" + string([]byte(os.Getenv("JWT_SECRET"))))
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return os.Getenv("JWT_SECRET"), nil
	})

	if claims, ok := token.Claims.(jwt.RegisteredClaims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return "", err
	}
}
