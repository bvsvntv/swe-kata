-- name: FetchCategories :many
SELECT
    *
FROM 
    category
ORDER BY category_id ASC
LIMIT $1
OFFSET $2;

-- name: CountCategories :one
SELECT
    COUNT(*)
FROM
    category;

-- name: GetCategory :one
SELECT
    *
FROM 
    category
WHERE category_id = $1
LIMIT 1;

-- name: CreateCategory :one
INSERT INTO
    category
(name, last_update)
VALUES
($1, NOW())
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM 
    category
WHERE category_id = $1;

-- name: UpdateCategory :one
UPDATE 
    category
SET 
    name = $2
WHERE category_id = $1
RETURNING *;

-- name: UpdateCategoryPartial :one
UPDATE
    category
SET
    name = COALESCE(sqlc.narg('name'), name)
WHERE category_id = sqlc.arg('category_id')
RETURNING *;

-- name: FetchCategoryFilms :many
SELECT
    film.*
FROM 
    film
JOIN 
    film_category 
    ON
    film_category.film_id = film.film_id
WHERE 
    film_category.category_id = $1
ORDER BY title ASC
LIMIT $2
OFFSET $3;
