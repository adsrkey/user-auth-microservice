-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA IF NOT EXISTS users;

CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS users.users
(
    id          uuid DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT pk_user
            PRIMARY KEY,
    created_at  timestamp,
    modified_at timestamp,
    email       varchar(250) UNIQUE            NOT NULL,
    hash        varchar(250)                   NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE users.users;

DROP SCHEMA users;

-- +goose StatementEnd
