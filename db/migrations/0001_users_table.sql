CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY,
    email varchar(64),
    password varchar(128)
);