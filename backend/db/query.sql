-- name: ListUsers :many
SELECT * FROM user
ORDER BY display_name;