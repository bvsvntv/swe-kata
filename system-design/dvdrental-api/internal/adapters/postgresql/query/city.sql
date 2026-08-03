-- name: FetchCities :many
SELECT
    *
FROM 
    city
LIMIT $1
OFFSET $2;

-- name: CountCities :one
SELECT
    COUNT(*)
FROM
    city;

-- name: GetCity :one
SELECT
    *
FROM 
    city
WHERE city_id = $1
LIMIT 1;

-- name: CreateCity :one
INSERT INTO
    city
(country_id, city, last_update)
VALUES
($1, $2, NOW())
RETURNING *;

-- name: DeleteCity :exec
DELETE FROM 
    city
WHERE city_id = $1;

-- name: UpdateCity :one
UPDATE 
    city
SET 
    country_id = $2,
    city = $3
WHERE city_id = $1
RETURNING *;

-- name: UpdateCityPartial :one
UPDATE
    city
SET
    city = COALESCE(sqlc.narg('city'), city)
WHERE city_id = sqlc.arg('city_id')
RETURNING *;
