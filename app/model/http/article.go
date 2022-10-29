package http

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type CreateArticleRequest struct {
	Title   string    `db:"title" json:"title" validate:"required"`
	Content string    `db:"content" json:"content" validate:"required"`
	Author  uuid.UUID `db:"author" json:"author" validate:"required"`
}

type PayloadJWT struct {
	jwt.StandardClaims
}

func (payload *PayloadJWT) Valid() error {
	if time.Now().After(time.Unix(payload.ExpiresAt, 0)) {
		return errors.New("expired")
	}
	return nil
}
