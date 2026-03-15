--- Queries for Login ---

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1
                      and deleted_at IS NULL;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

--- Queries for Register ---

-- name: CreateUser :exec
INSERT INTO users (id, username, email, password_hash) VALUES ($1, $2, $3, $4);