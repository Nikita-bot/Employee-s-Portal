-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN boss INT,
ADD CONSTRAINT fk_boss FOREIGN KEY (boss) REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
