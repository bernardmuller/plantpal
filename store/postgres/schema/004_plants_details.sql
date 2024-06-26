-- +goose Up
ALTER TABLE plants
ADD latin VARCHAR(255),
ADD category VARCHAR(255),
ADD origin VARCHAR(255),
ADD climate VARCHAR(255),
ADD tempMax VARCHAR(255),
ADD tempMin VARCHAR(255),
ADD idealLight VARCHAR(255),
ADD toleratedLight VARCHAR(255),
ADD watering VARCHAR(255),
ADD insects VARCHAR(255),
ADD diseases VARCHAR(255),
ADD soil VARCHAR(255),
ADD repotPeriod VARCHAR(255),
ADD use VARCHAR(255);

-- +goose Down
ALTER TABLE plants
DROP COLUMN latin,
DROP COLUMN category,
DROP COLUMN origin,
DROP COLUMN climate,
DROP COLUMN tempMax,
DROP COLUMN tempMin,
DROP COLUMN idealLight,
DROP COLUMN toleratedLight,
DROP COLUMN watering,
DROP COLUMN insects,
DROP COLUMN diseases,
DROP COLUMN soil,
DROP COLUMN repotPeriod,
DROP COLUMN use;
