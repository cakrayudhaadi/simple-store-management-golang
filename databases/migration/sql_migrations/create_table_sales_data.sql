-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS sales_data
(
    id SERIAL PRIMARY KEY,
    branch_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    employee_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    sold_date TIMESTAMP NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS sales_data;

-- +migrate StatementEnd