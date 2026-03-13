-- +goose Up
SELECT 'up SQL query';

CREATE TABLE if NOT EXISTS users (
    id UUID serial PRIMARY KEY,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
    password_hash text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
    );
-- +goose Down
SELECT 'down SQL query';

DROP TABLE IF EXISTS users;