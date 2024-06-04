-- +goose Up
ALTER TABLE sessions
ADD COLUMN ip_address TEXT NOT NULL;

-- +goose Down
ALTER TABLE sessions
DROP COLUMN ip_address;