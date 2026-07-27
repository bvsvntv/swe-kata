-- name: FetchActors :many
SELECT
    *
FROM 
    actor
LIMIT $1
OFFSET $2;

-- name: CountActors :one
SELECT
    COUNT(*)
FROM
    actor;
