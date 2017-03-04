
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE blogs ADD COLUMN is_authorize TINYINT(1) NOT NULL DEFAULT 0,
                  ADD COLUMN authorize_pass VARCHAR(100) DEFAULT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE blogs DROP is_authorize, authorize_pass;
