-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
)
VALUES (
    $1, $2, $3
) RETURNING id,username,email,created_at;

-- name: GetUser :one
SELECT * FROM users
where email =$1;
