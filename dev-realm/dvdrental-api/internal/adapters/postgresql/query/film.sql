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
