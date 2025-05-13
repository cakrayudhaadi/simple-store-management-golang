-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS employee
(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    branch_id INTEGER NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS employee;

-- +migrate StatementEnd