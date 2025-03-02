-- name: GetUser :one
SELECT
	*
FROM
	users
WHERE
	id = $1;

-- name: GetUserByEmail :one
SELECT
	*
FROM
	users
WHERE
	email = $1;

-- name: CreateUser :one
INSERT INTO
	users (email, password_hash)
VALUES ($1, $2) RETURNING *;