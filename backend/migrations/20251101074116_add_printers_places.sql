-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS printers (
    id SERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
