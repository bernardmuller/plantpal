-- +goose Up
ALTER TABLE users
ADD COLUMN provider TEXT;

-- +goose Down
ALTER TABLE users
DROP COLUMN provider;