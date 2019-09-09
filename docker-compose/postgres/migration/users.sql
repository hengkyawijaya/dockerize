DROP TABLE IF EXISTS users;

CREATE TABLE users (id serial  PRIMARY KEY,
name text  NOT NULL,
email text  NOT NULL
);

INSERT INTO users(name,email) VALUES('Ini Nama','hello@kid.con');