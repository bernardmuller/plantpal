-- name: CreatePlant :one
INSERT INTO plants (id, common, family, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: GetAllPlants :many
SELECT * FROM plants
ORDER BY created_at DESC;

-- name: UpdatePlant :one
UPDATE plants
SET common = $2, family = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePlant :one
DELETE FROM plants
WHERE id = $1
RETURNING *;

-- name: GetPlantByID :one
SELECT * FROM plants
WHERE id = $1;

-- name: GetPlantByCommon :one
SELECT * FROM plants
WHERE common = $1;

