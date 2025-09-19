-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS branches (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

INSERT INTO branches(name)
VALUES ('Островского'),('Александрова');

ALTER TABLE employee 
ADD COLUMN IF NOT EXISTS branch_id INT DEFAULT 1 REFERENCES branches(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
