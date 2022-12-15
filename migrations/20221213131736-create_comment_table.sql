-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS comment (
    id SERIAL,
    id_post BIGINT,
    id_user BIGINT,
    comment TEXT,
    created_at DATE,
    CONSTRAINT comment_id PRIMARY KEY (id),
    CONSTRAINT fk_post FOREIGN KEY (id_post) REFERENCES post(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES account(id) ON DELETE CASCADE
)
-- +migrate StatementEnd
-- +migrate Down
DROP TABLE comment;