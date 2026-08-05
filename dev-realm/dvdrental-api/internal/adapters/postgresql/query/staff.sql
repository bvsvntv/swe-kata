-- name: FetchStaffs :many
SELECT
    *
FROM
    staff
LIMIT $1
OFFSET $2;

-- name: CountStaffs :one
SELECT
    COUNT(*)
FROM
    staff;

-- name: GetStaff :one
SELECT
    *
FROM
    staff
WHERE staff_id = $1
LIMIT 1;

-- name: CreateStaff :one
INSERT INTO
    staff
(
    first_name,
    last_name,
    email, 
    address_id,
    store_id,
    active, 
    username,
    password, 
    picture,
    last_update
)
VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW()
)
RETURNING *;

-- name: UpdateStaff :one
UPDATE 
    staff
SET
    first_name = $2,
    last_name = $3,
    email = $4,
    address_id = $5,
    store_id = $6,
    active = $7,
    username = $8,
    password = $9,
    picture = $10,
    last_update = NOW()
WHERE staff_id = $1
RETURNING *;

-- name: UpdateStaffPartial :one
UPDATE 
    staff
SET
    first_name = COALESCE(sqlc.narg(first_name), first_name),
    last_name = COALESCE(sqlc.narg(last_name), last_name),
    email = COALESCE(sqlc.narg(email), email),
    address_id = COALESCE(sqlc.narg(address_id), address_id),
    store_id = COALESCE(sqlc.narg(store_id), store_id),
    active = COALESCE(sqlc.narg(active), active),
    username = COALESCE(sqlc.narg(username), username),
    password = COALESCE(sqlc.narg(password), password),
    picture = COALESCE(sqlc.narg(picture), picture),
    last_update = NOW()
WHERE staff_id = sqlc.arg(staff_id)
RETURNING *;

-- name: DeleteStaff :exec
DELETE FROM
    staff
WHERE staff_id = $1;

-- name: FetchStaffRentals :many
SELECT
    rental.*
FROM
    rental
WHERE
    rental.staff_id = $1
ORDER BY rental_date DESC
LIMIT $2
OFFSET $3;
