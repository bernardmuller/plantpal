-- +goose Up
CREATE TABLE plants
(
    id         UUID PRIMARY KEY,
    common     TEXT      NOT NULL,
    family     TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE plants;