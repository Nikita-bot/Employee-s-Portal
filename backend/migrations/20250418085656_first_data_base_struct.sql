-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS departments(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS users(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT '',
    surname VARCHAR DEFAULT '',
    patronymic VARCHAR DEFAULT '',
    position VARCHAR DEFAULT '',
    department_id INT REFERENCES departments(id),
    login VARCHAR DEFAULT '',
    password VARCHAR DEFAULT '',
    email VARCHAR DEFAULT '',
    phone VARCHAR DEFAULT '',
    tg_link VARCHAR DEFAULT '',
    tg_id BIGINT DEFAULT 0,
    pasport VARCHAR DEFAULT '',
    snyls VARCHAR DEFAULT '',
    adress VARCHAR DEFAULT '',
    employment_date VARCHAR DEFAULT '',
    dismissal_date VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS roles(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS tasks(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT '',
    department_id INT REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS user_task(
    id serial4 PRIMARY KEY,
    task_id INT REFERENCES tasks(id),
    initiator INT REFERENCES users(id),
    executor INT REFERENCES users(id),
    description VARCHAR DEFAULT '',
    status INT DEFAULT 0,
    create_date VARCHAR DEFAULT '',
    execute_date VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS user_role(
    role_id INT REFERENCES roles(id),
    user_id INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments(
    id serial4 PRIMARY KEY,
    user_task_id INT REFERENCES user_task(id),
    author_id INT REFERENCES users(id),
    comment VARCHAR DEFAULT '',
    creation_date VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS task_journal(
    user_task_id INT REFERENCES user_task(id),
    action VARCHAR DEFAULT '',
    creation_date VARCHAR DEFAULT ''
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
