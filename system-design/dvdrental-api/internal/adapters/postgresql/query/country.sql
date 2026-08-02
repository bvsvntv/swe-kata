-- name: FetchCountries :many
SELECT
    *
FROM 
    country
LIMIT $1
OFFSET $2;

-- name: CountCountries :one
SELECT
    COUNT(*)
FROM
    country;

-- name: GetCountry :one
SELECT
    *
FROM 
    country
WHERE country_id = $1
LIMIT 1;

-- name: CreateCountry :one
INSERT INTO
    country
(country, last_update)
VALUES
($1, NOW())
RETURNING *;

-- name: DeleteCountry :exec
DELETE FROM 
    country
WHERE country_id = $1;

-- name: UpdateCountry :one
UPDATE 
    country
SET 
    country = $2
WHERE country_id = $1
RETURNING *;
