
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE articles (
    id INT NOT NULL,
    blog_id INT NOT NULL,
    subject VARCHAR(255) NOT NULL,
    body TEXT,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    INDEX(blog_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE articles;
