-- +goose Up
-- +goose StatementBegin
ALTER TABLE departments ADD COLUMN IF NOT EXISTS is_die BOOLEAN DEFAULT FALSE;
ALTER TABLE roles ADD COLUMN IF NOT EXISTS description VARCHAR DEFAULT '';

CREATE TABLE IF NOT EXISTS employee (
    id VARCHAR(36) PRIMARY KEY,
    tab_num VARCHAR(50),
    zanyatost VARCHAR(100),
    start_date VARCHAR DEFAULT '',
    end_date VARCHAR DEFAULT '',
    position VARCHAR(255),
    depart_id INTEGER REFERENCES departments(id),
    user_id INTEGER REFERENCES users(id),
    boss INTEGER REFERENCES users(id)
);

ALTER TABLE users
    DROP COLUMN IF EXISTS position,
    DROP COLUMN IF EXISTS boss,
    DROP COLUMN IF EXISTS department_id,
    DROP COLUMN IF EXISTS employment_date,
    DROP COLUMN IF EXISTS dismissal_date,
    DROP COLUMN IF EXISTS pasport,
    ADD COLUMN IF NOT EXISTS pasport_ser VARCHAR DEFAULT '',
    ADD COLUMN IF NOT EXISTS pasport_num VARCHAR DEFAULT '',
    ADD COLUMN IF NOT EXISTS pasport_date VARCHAR DEFAULT '',
    ADD COLUMN IF NOT EXISTS pasport_dep VARCHAR DEFAULT '',
    ADD COLUMN IF NOT EXISTS pasport_dep_key VARCHAR DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE departments DROP COLUMN is_die;
ALTER TABLE roles DROP COLUMN description;
DROP TABLE employee;
ALTER TABLE users
    ADD COLUMN position VARCHAR,
    ADD COLUMN boss INTEGER,
    ADD COLUMN employment_date DATE,
    ADD COLUMN dismissal_date DATE,
    ADD COLUMN pasport VARCHAR,
    DROP COLUMN pasport_ser,
    DROP COLUMN pasport_num,
    DROP COLUMN pasport_date,
    DROP COLUMN pasport_dep,
    DROP COLUMN pasport_dep_key;
-- +goose StatementEnd