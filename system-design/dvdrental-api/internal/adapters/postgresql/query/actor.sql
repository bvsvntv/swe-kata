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

-- name: UpdateActor :one
UPDATE 
    actor
SET 
    first_name = $2,
    last_name = $3
WHERE actor_id = $1
RETURNING *;

-- name: UpdateActorPartial :one
UPDATE
    actor
SET
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name)
WHERE actor_id = sqlc.arg('actor_id')
RETURNING *;

-- name: FetchActorFilms :many
SELECT
    film.*
FROM 
    film
JOIN 
    film_actor 
    ON
    film_actor.film_id = film.film_id
WHERE 
    film_actor.actor_id = $1
ORDER BY 
    title
DESC
LIMIT $2
OFFSET $3;
