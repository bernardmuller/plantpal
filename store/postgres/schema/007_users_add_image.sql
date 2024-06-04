-- +goose Up
ALTER TABLE users
ADD COLUMN image TEXT;

-- +goose Down
ALTER TABLE users
DROP COLUMN image;