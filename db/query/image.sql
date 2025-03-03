-- name: GetUserImages :many
SELECT
	*
FROM
	images
WHERE
	user_id = $1
OFFSET
	$2
LIMIT
	$3;

-- name: GetImage :one
SELECT
	*
FROM
	images
WHERE
	id = $1;


-- name: CreateImage :one
INSERT INTO
	images (name, user_id, url)
VALUES
	($1, $2, $3)
RETURNING
	*;
