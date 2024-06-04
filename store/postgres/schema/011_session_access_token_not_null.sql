-- +goose Up
ALTER TABLE sessions
ALTER COLUMN access_token SET NOT NULL;

-- +goose Down
ALTER TABLE sessions
DROP COLUMN access_token;