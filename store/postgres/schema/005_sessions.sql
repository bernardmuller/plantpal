-- +goose Up
CREATE TABLE sessions
(
    id           UUID PRIMARY KEY,
    user_id      UUID REFERENCES users(id),
    access_token TEXT,
    expires      TIMESTAMP,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE sessions;