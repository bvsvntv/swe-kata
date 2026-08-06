-- name: FetchFilms :many
SELECT
    *
FROM 
    film
ORDER BY film_id ASC
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

-- name: FetchFilmActors :many
SELECT
    actor.*
FROM 
    actor
JOIN 
    film_actor 
    ON
    film_actor.actor_id = actor.actor_id
WHERE 
    film_actor.film_id = $1
ORDER BY first_name ASC
LIMIT $2
OFFSET $3;

-- name: FetchFilmCategories :many
SELECT
    category.*
FROM 
    category
JOIN 
    film_category
    ON
    film_category.category_id = category.category_id
WHERE 
    film_category.film_id = $1
ORDER BY name ASC
LIMIT $2
OFFSET $3;

-- name: FetchFilmInventory :many
SELECT
    inventory.*
FROM
    inventory
WHERE
    inventory.film_id = $1
ORDER BY film_id ASC
LIMIT $2
OFFSET $3;

-- name: CreateFilm :one
INSERT INTO film (
    title,
    description,
    release_year,
    language_id,
    rental_duration,
    rental_rate,
    length,
    replacement_cost,
    rating,
    special_features
) VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateFilm :one
UPDATE film
SET
    title = $2,
    description = $3,
    release_year = $4,
    language_id = $5,
    rental_duration = $6,
    rental_rate = $7,
    length = $8,
    replacement_cost = $9,
    rating = $10,
    special_features = $11,
    last_update = NOW()
WHERE film_id = $1
RETURNING *;

-- name: DeleteFilm :exec
DELETE FROM film
WHERE film_id = $1;
