// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: image.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createImage = `-- name: CreateImage :one
INSERT INTO
	images (name, user_id, url)
VALUES
	($1, $2, $3)
RETURNING
	id, name, user_id, url, uploaded_at
`

type CreateImageParams struct {
	Name   string      `json:"name"`
	UserID pgtype.UUID `json:"user_id"`
	Url    string      `json:"url"`
}

func (q *Queries) CreateImage(ctx context.Context, arg CreateImageParams) (Image, error) {
	row := q.db.QueryRow(ctx, createImage, arg.Name, arg.UserID, arg.Url)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Url,
		&i.UploadedAt,
	)
	return i, err
}

const getImage = `-- name: GetImage :one
SELECT
	id, name, user_id, url, uploaded_at
FROM
	images
WHERE
	id = $1
`

func (q *Queries) GetImage(ctx context.Context, id pgtype.UUID) (Image, error) {
	row := q.db.QueryRow(ctx, getImage, id)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Url,
		&i.UploadedAt,
	)
	return i, err
}

const getUserImages = `-- name: GetUserImages :many
SELECT
	id, name, user_id, url, uploaded_at
FROM
	images
WHERE
	user_id = $1
OFFSET
	$2
LIMIT
	$3
`

type GetUserImagesParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Offset int32       `json:"offset"`
	Limit  int32       `json:"limit"`
}

func (q *Queries) GetUserImages(ctx context.Context, arg GetUserImagesParams) ([]Image, error) {
	rows, err := q.db.Query(ctx, getUserImages, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Image{}
	for rows.Next() {
		var i Image
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Url,
			&i.UploadedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
