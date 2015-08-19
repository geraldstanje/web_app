CREATE TABLE IF NOT EXISTS account (
  email VARCHAR(32) NOT NULL PRIMARY KEY,
  username VARCHAR(32) NOT NULL UNIQUE COMMENT 'the username and password together comprise credentials required to obtain access (shortly login)',
  password CHAR(32) NOT NULL COMMENT 'the username and password together comprise credentials required to obtain access (shortly login)'
);