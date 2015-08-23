CREATE TABLE IF NOT EXISTS account (
  email VARCHAR(32) NOT NULL PRIMARY KEY,
  password VARCHAR(32) NOT NULL
);

INSERT INTO account VALUES('Christiano.Ronaldo@gmail.com', 'root');