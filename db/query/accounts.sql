-- name: GetAcount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;