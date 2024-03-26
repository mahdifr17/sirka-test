/**
  This is the SQL script that will be used to initialize the database schema.
  */

CREATE DATABASE test_case;

CREATE TABLE users (
  user_id VARCHAR (50) PRIMARY KEY,
  name VARCHAR (50) NOT NULL
);

INSERT INTO users (user_id, name) VALUES ('01', 'budi');
INSERT INTO users (user_id, name) VALUES ('02', 'Nano');
