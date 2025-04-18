// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: projects.sql

package sqlc

import (
	"context"
)

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, name, root_path, created_at FROM projects WHERE id = ?
`

func (q *Queries) GetProjectByID(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.RootPath,
		&i.CreatedAt,
	)
	return i, err
}

const insertProject = `-- name: InsertProject :exec
INSERT INTO projects (name,root_path) VALUES (?,?)
`

type InsertProjectParams struct {
	Name     string `json:"name"`
	RootPath string `json:"root_path"`
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) error {
	_, err := q.db.ExecContext(ctx, insertProject, arg.Name, arg.RootPath)
	return err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, root_path, created_at FROM projects ORDER BY created_at DESC
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.RootPath,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
