-- name: CreateUser :one
INSERT INTO users (
  full_name, email, password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 and deleted_at is null;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 and deleted_at is null;

-- name: GetLastUser :one
SELECT * FROM users where deleted_at is null ORDER BY id desc LIMIT 1;

-- name: ExistsAnUserUsingTheSameEmail :one
SELECT count(1) > 0 as exists FROM users WHERE email = $1 and deleted_at is null;