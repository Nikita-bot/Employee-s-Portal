-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS departments(
    id serial4 PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS users(
    id serial4 PRIMARY KEY,
    name VARCHAR,
    surname VARCHAR,
    patronymic VARCHAR,
    position VARCHAR,
    department_id INT REFERENCES departments(id),
    login VARCHAR,
    password VARCHAR,
    email VARCHAR,
    phone VARCHAR,
    tg_link VARCHAR,
    tg_id BIGINT,
    pasport VARCHAR,
    snyls VARCHAR,
    adress VARCHAR,
    employment_date DATE,
    dismissal_date DATE
);

CREATE TABLE IF NOT EXISTS roles(
    id serial4 PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS tasks(
    id serial4 PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS task_dep(
    task_id INT REFERENCES tasks(id),
    department_id INT REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS user_task(
    id serial4 PRIMARY KEY,
    task_id INT REFERENCES tasks(id),
    initiator INT REFERENCES users(id),
    executor INT REFERENCES users(id),
    description VARCHAR,
    status INT,
    create_date DATE,
    execute_date DATE
);

CREATE TABLE IF NOT EXISTS user_role(
    role_id INT REFERENCES roles(id),
    user_id INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments(
    id serial4 PRIMARY KEY,
    user_task_id INT REFERENCES user_task(id),
    author_id INT REFERENCES users(id),
    comment VARCHAR,
    creation_date DATE
);

CREATE TABLE IF NOT EXISTS task_journal(
    user_task_id INT REFERENCES user_task(id),
    action VARCHAR,
    creation_date DATE
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
