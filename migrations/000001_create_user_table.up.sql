-- +migrate Up
-- SQL in the section 'Up' is executed when this migration is applied
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name TEXT NOT NULL
);

