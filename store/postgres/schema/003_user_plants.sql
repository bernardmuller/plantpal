-- +goose Up
CREATE TABLE user_plants
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    plant_id   UUID REFERENCES plants (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE user_plants;