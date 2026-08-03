-- name: FetchLanguages :many
SELECT
    *
FROM 
    language
ORDER BY name ASC
LIMIT $1
OFFSET $2;

-- name: CountLanguages :one
SELECT
    COUNT(*)
FROM
    language;

-- name: GetLanguage :one
SELECT
    *
FROM 
    language
WHERE language_id = $1
LIMIT 1;
