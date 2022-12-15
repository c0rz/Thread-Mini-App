-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id SERIAL,
    name VARCHAR(100),
    created_at DATE,
    CONSTRAINT category_id PRIMARY KEY (id)
)
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE category;