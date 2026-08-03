-- name: FetchAddresses :many
SELECT * FROM address
ORDER BY address_id
LIMIT $1 OFFSET $2;

-- name: CountAddresses :one
SELECT count(*) FROM address;

-- name: GetAddress :one
SELECT * FROM address
WHERE address_id = $1 LIMIT 1;

-- name: CreateAddress :one
INSERT INTO address (
    address,
    address2,
    district,
    city_id,
    postal_code,
    phone
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateAddress :one
UPDATE address
SET
    address = $2,
    address2 = $3,
    district = $4,
    city_id = $5,
    postal_code = $6,
    phone = $7,
    last_update = now()
WHERE address_id = $1
RETURNING *;

-- name: UpdateAddressPartial :one
UPDATE address
SET
    address = COALESCE(sqlc.narg('address'), address),
    address2 = COALESCE(sqlc.narg('address2'), address2),
    district = COALESCE(sqlc.narg('district'), district),
    city_id = COALESCE(sqlc.narg('city_id'), city_id),
    postal_code = COALESCE(sqlc.narg('postal_code'), postal_code),
    phone = COALESCE(sqlc.narg('phone'), phone),
    last_update = now()
WHERE address_id = sqlc.arg('address_id')
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE address_id = $1;
