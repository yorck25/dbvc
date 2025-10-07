CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(128),
    email      VARCHAR(128) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active     bool      DEFAULT true
);

CREATE TABLE user_login
(
    user_id         INT REFERENCES users (id),
    username        VARCHAR(50) UNIQUE NOT NULL,
    password_hash   VARCHAR(255)       NOT NULL,
    last_login_at   TIMESTAMP,
    failed_attempts INT
);


CREATE TABLE connection_types
(
    id          SERIAL PRIMARY KEY,
    type_name   VARCHAR(50) UNIQUE NOT NULL,
    key         VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    active      bool DEFAULT true
);

CREATE TABLE projects
(
    id              SERIAL PRIMARY KEY,
    owner_id        INT REFERENCES users (id),
    name            VARCHAR(100) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active          bool      DEFAULT true,
    connection_type INT REFERENCES connection_types (id)
);

CREATE TABLE versions
(
    id         SERIAL PRIMARY KEY,
    version    varchar(50) UNIQUE NOT NULL,
    up         jsonb,
    down       jsonb,
    state      varchar(128)
        CONSTRAINT state_check CHECK (state IN ('pending', 'completed', 'failed')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    applied_at TIMESTAMP,
    project_id INT REFERENCES projects (id)
);

CREATE TABLE version_audit
(
    id         SERIAL PRIMARY KEY,
    version_id INT REFERENCES versions (id),
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes      varchar(512)
);

CREATE TABLE user_role
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id),
    project_id INT REFERENCES projects (id),
    role varchar(255)
);

INSERT INTO connection_types (type_name, key, description)
VALUES ('PostgreSQL', 'psql', 'PostgreSQL Database Connection');

INSERT INTO users (first_name, email) VALUES ('Yorck', 'Dombrowsky');