-- name: CreateTask :one
INSERT INTO tasks (title, description, completed)
VALUES ($1, $2, $3)
RETURNING id, title, description, completed;

-- name: GetTask :one
SELECT id, title, description, completed
FROM tasks
WHERE id = $1;

-- name: ListTasks :many
SELECT id, title, description, completed
FROM tasks
ORDER BY id ASC;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2,
    description = $3,
    completed = $4
WHERE id = $1
RETURNING id, title, description, completed;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;