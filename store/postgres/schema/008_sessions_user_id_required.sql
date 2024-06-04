-- +goose Up
ALTER TABLE sessions
ALTER COLUMN user_id SET NOT NULL;

-- +goose Down
ALTER TABLE sessions
ALTER COLUMN user_id DROP NOT NULL;