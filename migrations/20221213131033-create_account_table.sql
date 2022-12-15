-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id SERIAL,
    full_name VARCHAR(100),
    email VARCHAR(50),
    password TEXT,
    level VARCHAR(10),
    created_at DATE,
    CONSTRAINT account_id PRIMARY KEY (id)
)
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE account;