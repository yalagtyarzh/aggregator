CREATE TABLE rating
(
    id       SERIAL PRIMARY KEY NOT NULL,
    name     VARCHAR(2)         NOT NULL,
    name_ext VARCHAR(25)        NOT NULL
);

CREATE TABLE products
(
    id           SERIAL PRIMARY KEY NOT NULL,
    title        TEXT               NOT NULL,
    description  TEXT               NOT NULL,
    year         INTEGER            NOT NULL,
    release_date TIMESTAMP          NOT NULL,
    studio       TEXT               NOT NULL,
    rating_id    INTEGER            NOT NULL REFERENCES rating (id),
    created_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP          NOT NULL DEFAULT NOW()
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
    role       TEXT NOT NULL REFERENCES roles (name),
    permission TEXT NOT NULL REFERENCES permissions (name),
    UNIQUE (role, permission)
);

CREATE TABLE users
(
    id         UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    first_name TEXT             NOT NULL,
    last_name  TEXT             NOT NULL,
    user_name  TEXT             NOT NULL UNIQUE,
    email      TEXT             NOT NULL UNIQUE,
    password   TEXT             NOT NULL,
    role       TEXT             NOT NULL REFERENCES roles (name),
    created_at TIMESTAMP        NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP        NOT NULL DEFAULT NOW()

);

CREATE TABLE reviews
(
    id           SERIAL PRIMARY KEY NOT NULL,
    score        INTEGER            NOT NULL,
    content      TEXT               NOT NULL,
    content_html TEXT               NOT NULL,
    created_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    CHECK (score BETWEEN 0 AND 100)
);

CREATE TABLE products_reviews
(
    product_id INTEGER NOT NULL REFERENCES products (id),
    review_id  INTEGER NOT NULL REFERENCES reviews (id)
);