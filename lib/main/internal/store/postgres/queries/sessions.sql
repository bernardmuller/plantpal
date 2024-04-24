-- name: CreateSession :one
INSERT INTO sessions (id, user_id, expires, access_token, ip_address )
VALUES ($1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetSessionByID :one
SELECT * FROM sessions WHERE id = $1;

-- name: DeleteSessionByID :exec
DELETE FROM sessions WHERE id = $1;
