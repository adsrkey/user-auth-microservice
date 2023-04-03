-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA development;

CREATE EXTENSION pgcrypto;

CREATE TABLE development.user (
    id           uuid         DEFAULT gen_random_uuid()      NOT NULL
    CONSTRAINT pk_user
    PRIMARY KEY,
    created_at   timestamp,
    modified_at  timestamp,
    email        varchar(250) UNIQUE NOT NULL,
    hash         varchar(250) NOT NULL
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE development.user;

DROP SCHEMA development;

-- +goose StatementEnd
