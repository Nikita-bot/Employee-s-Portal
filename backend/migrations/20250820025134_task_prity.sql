-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_task ADD COLUMN IF NOT EXISTS priority INTEGER DEFAULT 2;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
