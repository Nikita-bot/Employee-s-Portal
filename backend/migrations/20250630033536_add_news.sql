-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS news(
    id SERIAL PRIMARY KEY,
    title VARCHAR DEFAULT '',
    author INT REFERENCES users(id),
    content VARCHAR DEFAULT '',
    date VARCHAR DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
