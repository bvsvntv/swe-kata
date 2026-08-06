-- name: FetchRentals :many
SELECT
    *
FROM
    rental
ORDER BY rental_id ASC
LIMIT $1
OFFSET $2;

-- name: CountRentals :one
SELECT
    COUNT(*)
FROM
    rental;

-- name: GetRental :one
SELECT
    *
FROM
    rental
WHERE
    rental_id = $1
LIMIT 1;

-- name: CreateRental :one
INSERT INTO rental (
    rental_date,
    inventory_id,
    customer_id,
    staff_id,
    return_date
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: UpdateRental :one
UPDATE
    rental
SET
    rental_date = $2,
    inventory_id = $3,
    customer_id = $4,
    staff_id = $5,
    return_date = $6,
    last_update = NOW()
WHERE
    rental_id = $1
RETURNING *;

-- name: DeleteRental :exec
DELETE FROM
    rental
WHERE
    rental_id = $1;

-- name: ReturnRental :one
UPDATE
    rental
SET
    return_date = NOW(),
    last_update = NOW()
WHERE
    rental_id = $1
RETURNING *;
