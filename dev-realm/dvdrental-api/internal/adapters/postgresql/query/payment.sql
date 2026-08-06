-- name: FetchPayments :many
SELECT
    *
FROM
    payment
ORDER BY payment_id ASC
LIMIT $1
OFFSET $2;

-- name: CountPayments :one
SELECT
    COUNT(*)
FROM
    payment;

-- name: GetPayment :one
SELECT
    *
FROM 
    payment
WHERE payment_id = $1;

-- name: CreatePayment :one
INSERT INTO
    payment
(
    customer_id,
    rental_id,
    staff_id,
    amount,
    payment_date
)
VALUES
($1, $2, $3, $4, NOW())
RETURNING *;

-- name: UpdatePayment :one
UPDATE 
    payment
SET
    customer_id = $2,
    rental_id = $3,
    staff_id = $4,
    amount = $5,
    payment_date = NOW()
WHERE payment_id = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM 
    payment
WHERE payment_id = $1;
