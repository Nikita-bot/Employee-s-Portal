-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS task_roles (
    task_id INTEGER REFERENCES tasks(id),
    role_id INTEGER REFERENCES roles(id)
);

ALTER TABLE tasks RENAME TO task_list;

ALTER TABLE user_task RENAME TO tasks;

DROP TABLE IF EXISTS task_department;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
