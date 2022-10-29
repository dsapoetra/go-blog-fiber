package http

import (
	"time"
)

type CreateAuthorRequest struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FullName  string    `json:"full_name" validate:"required"`
	UserName  string    `json:"user_name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
}

type LoginAuthorRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginAuthorResponse struct {
	Token string `json:"token"`
}
