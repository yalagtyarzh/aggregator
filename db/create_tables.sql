CREATE TABLE rating
(
    id       SERIAL PRIMARY KEY NOT NULL,
    name     VARCHAR(2)         NOT NULL,
    name_ext VARCHAR(25)        NOT NULL,
);

CREATE TABLE products
(
    id           SERIAL PRIMARY KEY NOT NULL,
    title        TEXT               NOT NULL,
    description  TEXT               NOT NULL,
    year         INTEGER            NOT NULL,
    release_date TIMESTAMP          NOT NULL,
    studio       TEXT               NOT NULL,
    rating_id    INTEGER            NOT NULL REFERENCES rating (id)
);

CREATE TABLE roles
(
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE permissions
(
    name text NOT NULL UNIQUE
);

CREATE TABLE roles_permissions
(
    role        TEXT NOT NULL REFERENCES roles (name),
    permissions TEXT NOT NULL REFERENCES roles (permissions),
    UNIQUE (role, permissions)
);

CREATE TABLE users
(
    id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    first_name TEXT             NOT NULL,
    last_name  TEXT             NOT NULL,
    user_name  TEXT             NOT NULL UNIQUE,
    email      TEXT             NOT NULL UNIQUE,
    password   TEXT             NOT NULL,
    role       TEXT             NOT NULL REFERENCES roles (name)
);

CREATE TABLE reviews
(
    id           UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    score        INTEGER          NOT NULL,
    content      TEXT             NOT NULL,
    content_html TEXT             NOT NULL,
    CHECK (score BETWEEN 0 AND 100)
);