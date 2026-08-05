-- name: FetchStores :many
SELECT
    *
FROM 
    store
LIMIT $1
OFFSET $2;

-- name: CountStores :one
SELECT
    COUNT(*)
FROM
    store;

-- name: GetStore :one
SELECT
    *
FROM 
    store
WHERE store_id = $1
LIMIT 1;

-- name: CreateStore :one
INSERT INTO
    store
(
    manager_staff_id,
    address_id,
    last_update
)
VALUES
(
    $1,
    $2,
    NOW()
)
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM 
    store
WHERE store_id = $1;

-- name: UpdateStore :one
UPDATE
    store
SET
    manager_staff_id = $2,
    address_id = $3
WHERE store_id = $1
RETURNING *;

-- name: UpdateStorePartial :one
UPDATE
    store
SET
    manager_staff_id = COALESCE(sqlc.narg('manager_staff_id'), manager_staff_id),
    address_id = COALESCE(sqlc.narg('address_id'), address_id)
WHERE store_id = sqlc.arg('store_id')
RETURNING *;

-- name: FetchStoreInventory :many
SELECT
    inventory.*
FROM
    inventory
WHERE
    inventory.store_id = $1
ORDER BY store_id ASC
LIMIT $2
OFFSET $3;

-- name: FetchStoreCustomers :many
SELECT
    customer.*
FROM
    customer
WHERE
    customer.store_id = $1
ORDER BY store_id ASC
LIMIT $2
OFFSET $3;
