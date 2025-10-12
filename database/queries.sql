-- name: CreateTask :one
insert into tasks (name, description, due_at)
values (?, ?, ?)
returning task_id,
  name,
  description,
  due_at,
  created_at;

-- name: ListTasks :many
select task_id,
  name,
  description,
  due_at,
  created_at
from tasks;

-- name: UpdateTask :one
update tasks
set name = ?,
  description = ?,
  due_at = ?
where task_id = ?
returning task_id,
  name,
  description,
  due_at,
  created_at;
