-- name: FetchInventories :many
SELECT
    *
FROM 
    inventory
ORDER BY inventory_id ASC
LIMIT $1
OFFSET $2;

-- name: CountInventories :one
SELECT
    COUNT(*)
FROM
    inventory;

-- name: GetInventory :one
SELECT
    *
FROM 
    inventory
WHERE inventory_id = $1
LIMIT 1;

-- name: CreateInventory :one
INSERT INTO
    inventory
(
    film_id,
    store_id,
    last_update
)
VALUES
(
    $1,
    $2,
    NOW()
)
RETURNING *;

-- name: DeleteInventory :exec
DELETE FROM 
    inventory
WHERE inventory_id = $1;

-- name: FetchInventoryRentals :many
SELECT
    rental.*
FROM
    rental
WHERE
    rental.inventory_id = $1
ORDER BY rental_date DESC
LIMIT $2
OFFSET $3;
