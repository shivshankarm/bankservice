-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
where id = $1
Returning *;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE id = $1;



-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries
SET amount = $2
where account_id = $1
Returning *;

-- name: DeleteEntry :exec
DELETE FROM entries 
WHERE id = $1;


-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE from_account_id = $1 
AND to_account_id = $2
LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $3
WHERE from_account_id = $1 
AND to_account_id = $2
Returning *;

-- name: DeleteTransfer :exec
DELETE FROM transfers 
WHERE id = $1 ;