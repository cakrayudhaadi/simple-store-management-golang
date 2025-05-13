-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS item_type
(
    id SERIAL PRIMARY KEY,
    type VARCHAR NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS item_type;

-- +migrate StatementEnd