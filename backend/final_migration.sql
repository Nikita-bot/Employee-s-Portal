-- +goose Up
-- +goose StatementBegin

-- Создание таблиц departments и users
CREATE TABLE IF NOT EXISTS departments(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT '',
    is_die BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS users(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT '',
    surname VARCHAR DEFAULT '',
    patronymic VARCHAR DEFAULT '',
    login VARCHAR DEFAULT '',
    password VARCHAR DEFAULT '',
    email VARCHAR DEFAULT '',
    phone VARCHAR DEFAULT '',
    tg_link VARCHAR DEFAULT '',
    tg_id BIGINT DEFAULT 0,
    snyls VARCHAR DEFAULT '',
    adress VARCHAR DEFAULT '',
    pasport_ser VARCHAR DEFAULT '',
    pasport_num VARCHAR DEFAULT '',
    pasport_date VARCHAR DEFAULT '',
    pasport_dep VARCHAR DEFAULT '',
    pasport_dep_key VARCHAR DEFAULT ''
);

-- Добавление колонки boss в таблицу users
ALTER TABLE users
ADD COLUMN boss INT,
ADD CONSTRAINT fk_boss FOREIGN KEY (boss) REFERENCES users(id);

CREATE TABLE IF NOT EXISTS roles(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT '',
    description VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS task_list(
    id serial4 PRIMARY KEY,
    name VARCHAR DEFAULT '',
    type VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS tasks(
    id serial4 PRIMARY KEY,
    task_id INT REFERENCES task_list(id),
    initiator INT REFERENCES users(id),
    executor INT REFERENCES users(id),
    description VARCHAR DEFAULT '',
    status INT DEFAULT 0,
    create_date VARCHAR DEFAULT '',
    execute_date VARCHAR DEFAULT '',
    priority INTEGER DEFAULT 2
);

CREATE TABLE IF NOT EXISTS user_role(
    role_id INT REFERENCES roles(id),
    user_id INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments(
    id serial4 PRIMARY KEY,
    user_task_id INT REFERENCES tasks(id),
    author_id INT REFERENCES users(id),
    comment VARCHAR DEFAULT '',
    creation_date VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS task_journal(
    user_task_id INT REFERENCES tasks(id),
    action VARCHAR DEFAULT '',
    creation_date VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS news(
    id SERIAL PRIMARY KEY,
    title VARCHAR DEFAULT '',
    author INT REFERENCES users(id),
    content VARCHAR DEFAULT '',
    date VARCHAR DEFAULT ''
);

CREATE TABLE IF NOT EXISTS branches (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS employee (
    id VARCHAR(36) PRIMARY KEY,
    tab_num VARCHAR(50),
    zanyatost VARCHAR(100),
    start_date VARCHAR DEFAULT '',
    end_date VARCHAR DEFAULT '',
    position VARCHAR(255),
    depart_id INTEGER REFERENCES departments(id),
    branch_id INTEGER DEFAULT 1 REFERENCES branches(id),
    user_id INTEGER REFERENCES users(id),
    boss INTEGER REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS task_roles (
    task_id INTEGER REFERENCES task_list(id),
    role_id INTEGER REFERENCES roles(id)
);

-- +goose StatementEnd