-- name: GetProjectByID :one
SELECT * FROM projects WHERE id = ?;

-- name: ListProjects :many
SELECT * FROM projects ORDER BY created_at DESC;

-- name: InsertProject :exec
INSERT INTO projects (name,root_path) VALUES (?,?);