-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS post (
    id SERIAL,
    title VARCHAR(100),
    id_category BIGINT,
    text TEXT,
    id_user BIGINT,
    created_at DATE,
    CONSTRAINT post_id PRIMARY KEY (id),
    CONSTRAINT fk_category FOREIGN KEY (id_category) REFERENCES category(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES account(id) ON DELETE CASCADE
)
-- +migrate StatementEnd
-- +migrate Down
DROP TABLE post;