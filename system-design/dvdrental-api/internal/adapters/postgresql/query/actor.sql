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

-- name: GetActor :one
SELECT
    *
FROM 
    actor
WHERE actor_id = $1
LIMIT 1;

-- name: CreateActor :one
INSERT INTO
    actor
(first_name, last_name, last_update)
VALUES
($1, $2, NOW())
RETURNING *;

-- name: DeleteActor :exec
DELETE FROM 
    actor
WHERE actor_id = $1;
