DROP TABLE IF EXISTS ingredients;
CREATE TYPE unit_name AS ENUM ('gram', 'short');
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    unit_name unit_name
);