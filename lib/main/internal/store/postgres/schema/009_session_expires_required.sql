-- +goose Up
ALTER TABLE sessions
ALTER COLUMN expires SET NOT NULL;

-- +goose Down
ALTER TABLE sessions
ALTER COLUMN expires DROP NOT NULL;