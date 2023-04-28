CREATE TABLE rating
(
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE genres
(
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE products
(
    id           SERIAL PRIMARY KEY NOT NULL,
    title        TEXT               NOT NULL,
    description  TEXT               NOT NULL,
    year         INTEGER            NOT NULL,
    release_date TIMESTAMP          NOT NULL,
    studio       TEXT               NOT NULL,
    rating       TEXT               NOT NULL REFERENCES rating (name),
    created_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    is_deleted   BOOL               NOT NULL DEFAULT FALSE
);

CREATE TABLE products_genres
(
    product_id INTEGER NOT NULL REFERENCES products (id),
    genre      TEXT    NOT NULL REFERENCES genres (name)
);

CREATE TABLE roles
(
    name TEXT NOT NULL UNIQUE
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
    updated_at TIMESTAMP        NOT NULL DEFAULT NOW(),
    is_deleted BOOL             NOT NULL DEFAULT FALSE
);

CREATE TABLE reviews
(
    id           SERIAL PRIMARY KEY NOT NULL,
    score        INTEGER            NOT NULL,
    content      TEXT               NOT NULL,
    content_html TEXT               NOT NULL,
    user_id      UUID               NOT NULL REFERENCES users (id),
    created_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP          NOT NULL DEFAULT NOW(),
    CHECK (score BETWEEN 0 AND 100)
);

CREATE TABLE products_reviews
(
    product_id INTEGER NOT NULL REFERENCES products (id),
    review_id  INTEGER NOT NULL REFERENCES reviews (id)
);

CREATE TABLE users_tokens
(
    user_id       UUID NOT NULL REFERENCES users (id) UNIQUE,
    refresh_token TEXT NOT NULL
);