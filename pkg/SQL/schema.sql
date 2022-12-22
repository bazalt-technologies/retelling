CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    age INTEGER NOT NULL DEFAULT 0,
    review_count INTEGER NOT NULL DEFAULT 0,
    rating INTEGER NOT NULL DEFAULT 0,
    profession TEXT NOT NULL DEFAULT '',
    likes INTEGER[] DEFAULT array[]::integer[]
);

CREATE TABLE types(
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL DEFAULT ''
);

CREATE TABLE genres(
    id SERIAL PRIMARY KEY,
    genre TEXT NOT NULL DEFAULT ''
);

CREATE TABLE content (
    id SERIAL PRIMARY KEY,
    type_id INTEGER NOT NULL REFERENCES types(id),
    genre1_id INTEGER NOT NULL REFERENCES genres(id),
    genre2_id INTEGER REFERENCES genres(id),
    genre3_id INTEGER REFERENCES genres(id),
    title TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    users_liked INTEGER[] DEFAULT NULL
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    content_id INTEGER NOT NULL REFERENCES content(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    review TEXT DEFAULT '',
    date INTEGER NOT NULL DEFAULT 0
);

