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

-- name: GetFilm :one
SELECT
    *
FROM 
    film
WHERE film_id = $1
LIMIT 1;
