// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Image struct {
	ID         pgtype.UUID        `json:"id"`
	Name       string             `json:"name"`
	UserID     pgtype.UUID        `json:"user_id"`
	Url        string             `json:"url"`
	UploadedAt pgtype.Timestamptz `json:"uploaded_at"`
}

type ImageTransformation struct {
	ID             pgtype.UUID        `json:"id"`
	ImageID        pgtype.UUID        `json:"image_id"`
	Transformation []byte             `json:"transformation"`
	Url            string             `json:"url"`
	TransformedAt  pgtype.Timestamptz `json:"transformed_at"`
}

type User struct {
	ID           pgtype.UUID        `json:"id"`
	Email        string             `json:"email"`
	PasswordHash string             `json:"password_hash"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}
