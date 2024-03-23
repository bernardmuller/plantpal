-- name: CreatePlant :one
INSERT INTO plants (id, common, family, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: GetAllPlants :many
SELECT * FROM plants
ORDER BY created_at DESC;