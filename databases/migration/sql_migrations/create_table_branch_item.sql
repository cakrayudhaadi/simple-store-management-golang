-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS branch_item
(
    id SERIAL PRIMARY KEY,
    branch_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS branch_item;

-- +migrate StatementEnd