
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE blogs ADD COLUMN blog_id VARCHAR(25) NOT NULL AFTER id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE blogs DROP blog_id;
