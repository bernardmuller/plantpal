-- name: CreatePlant :one
INSERT INTO plants (id, common, family, created_at, updated_at, latin, category, origin, climate, tempmax, tempmin, ideallight, toleratedlight, watering, insects, diseases, soil, repotperiod, use)
VALUES ($1, $2, $3, NOW(), NOW(), $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
RETURNING *;

-- name: GetAllPlants :many
SELECT * FROM plants
ORDER BY created_at DESC;

-- name: UpdatePlant :one
UPDATE plants
SET common = $2, family = $3, updated_at = NOW(), latin = $4, category = $5, origin = $6, climate = $7, tempmax = $8, tempmin = $9, ideallight = $10, toleratedlight = $11, watering = $12, insects = $13, diseases = $14, soil = $15, repotperiod = $16, use = $17
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

