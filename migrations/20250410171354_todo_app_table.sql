-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    email text not null UNIQUE,
    pass_hash bytea not null
);
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE,
    uid int NOT NULL,
    foreign key(uid) references users(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS goose_db_version (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN NOT NULL,
    tstamp TIMESTAMP DEFAULT NOW()
);
-- +goose Down
DROP TABLE tasks;
DROP TABLE users;

