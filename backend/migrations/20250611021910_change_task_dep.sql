-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS task_department(
    task_id INT REFERENCES tasks(id),
    department_id INT REFERENCES departments(id)
);

INSERT INTO task_department (task_id, department_id)
SELECT id, department_id FROM tasks WHERE department_id IS NOT NULL;

ALTER TABLE tasks
DROP COLUMN department_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
