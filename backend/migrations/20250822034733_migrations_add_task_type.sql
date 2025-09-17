-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks ADD COLUMN IF NOT EXISTS type VARCHAR DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
