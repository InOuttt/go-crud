CREATE TABLE user
(
    id         VARCHAR PRIMARY KEY,
    username   VARCHAR NOT NULL,
    password   VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE profile
(
    id      VARCHAR PRIMARY KEY,
    user_id  VARCHAR NOT NULL,
    email   VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
)
