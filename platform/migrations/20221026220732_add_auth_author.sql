-- +goose Up
-- +goose StatementBegin
ALTER TABLE authors
    ADD COLUMN password VARCHAR(500),
    ADD COLUMN username VARCHAR(500),
    ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE authors
    DROP COLUMN password,
    DROP COLUMN username,
    DROP COLUMN is_deleted;
-- +goose StatementEnd
