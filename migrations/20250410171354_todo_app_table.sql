-- +goose Up
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE
);
-- CREATE TABLE IF NOT EXISTS goose_db_version (
--     id SERIAL PRIMARY KEY,
--     version_id BIGINT NOT NULL,
--     is_applied BOOLEAN NOT NULL,
--     tstamp TIMESTAMP DEFAULT NOW()
-- );
-- +goose Down
DROP TABLE tasks;

