package model

import (
	"github.com/google/uuid"
	"time"
)

type Author struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	FullName  string    `db:"full_name" json:"full_name" validate:"required"`
}