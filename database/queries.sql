-- name: CreateTask :one
insert into tasks (name, description, due_at)
values (?, ?, ?)
returning *;
