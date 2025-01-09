-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: CreateUser :exec
INSERT INTO users (
    name, email, age
) VALUES (
    $1, $2, $3
);

-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    email = $3,
    age = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;