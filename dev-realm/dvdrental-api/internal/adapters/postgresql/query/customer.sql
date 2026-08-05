-- name: FetchCustomers :many
SELECT
    *
FROM
    customer
LIMIT $1
OFFSET $2;

-- name: CountCustomers :one
SELECT
    COUNT(*)
FROM
    customer;

-- name: GetCustomer :one
SELECT
    *
FROM
    customer
WHERE
    customer_id = $1;

-- name: CreateCustomer :one
INSERT INTO customer (
    first_name,
    last_name,
    email,
    address_id,
    store_id,
    activebool,
    active
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    CASE
        WHEN $6 THEN 1
        ELSE 0
    END
)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customer
SET
    first_name = $1,
    last_name = $2,
    email = $3,
    address_id = $4,
    store_id = $5,
    activebool = $6,
    active = CASE
        WHEN $6 THEN 1
        ELSE 0
    END,
    last_update = NOW()
WHERE
    customer_id = $7
RETURNING *;

-- name: UpdateCustomerPartial :one
UPDATE customer
SET
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name),
    email = COALESCE(sqlc.narg('email'), email),
    address_id = COALESCE(sqlc.narg('address_id'), address_id),
    store_id = COALESCE(sqlc.narg('store_id'), store_id),
    activebool = COALESCE(sqlc.narg('activebool'), activebool),
    active = CASE
        WHEN sqlc.narg('activebool') IS NULL THEN active
        WHEN sqlc.narg('activebool') THEN 1
        ELSE 0
    END,
    last_update = NOW()
WHERE
    customer_id = sqlc.arg('customer_id')
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM
    customer
WHERE
    customer_id = $1;

-- name: FetchCustomerRentals :many
SELECT
    rental.*
FROM
    rental
WHERE
    rental.customer_id = $1
ORDER BY rental_date DESC
LIMIT $2
OFFSET $3;
