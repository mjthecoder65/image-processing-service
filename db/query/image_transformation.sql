-- name: CreateImageTransformation :one
INSERT INTO image_transformations(
    image_id,
    transformation,
    url
)
VALUES(
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetImageTransformations :many
SELECT 
    * 
FROM 
    image_transformations
WHERE 
    image_id = $1
OFFSET $2 
LIMIT $3;