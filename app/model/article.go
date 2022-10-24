package model

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Title     string    `db:"title" json:"title" validate:"required"`
	Content   string    `db:"content" json:"content" validate:"required"`
	Author    uuid.UUID `db:"author" json:"author" validate:"required"`
	Likes     uint64    `db:"likes" json:"likes"`
	IsDeleted bool      `db:"is_deleted" json:"is_deleted"`
}
