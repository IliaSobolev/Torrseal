CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    lang CHAR(5) NOT NULL,
    expiry_date TIMESTAMPTZ
);