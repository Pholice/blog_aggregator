-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE,
    feed_id UUID NOT NULL,
    description TEXT NULL, 
    published_at TIMESTAMP NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id)
);

-- +goose Down 
DROP TABLE posts;