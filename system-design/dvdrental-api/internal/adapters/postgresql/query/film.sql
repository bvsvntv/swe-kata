-- name: FetchFilms :many
SELECT
    *
FROM 
    film
LIMIT $1
OFFSET $2;

-- name: CountFilms :one
SELECT
    COUNT(*)
FROM
    film;
