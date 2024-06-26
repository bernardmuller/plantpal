-- +goose Up
CREATE TABLE users
(
    id         UUID PRIMARY KEY,
    email      TEXT      NOT NULL,
    firstname  TEXT      NOT NULL,
    lastname   TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;