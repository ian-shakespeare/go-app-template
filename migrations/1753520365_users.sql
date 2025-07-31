create table if not exists users (
  user_id integer primary key,
  email text not null unique check (5 <= length(email) and length(email) <= 64)
);
