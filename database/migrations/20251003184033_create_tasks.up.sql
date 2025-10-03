create table if not exists tasks (
  task_id integer primary key,
  name text not null check(length(name) <= 64),
  description text,
  due_at datetime,
  created_at datetime not null default current_timestamp
);
