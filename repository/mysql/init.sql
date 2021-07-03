CREATE USER 'myuser' IDENTIFIED BY 'password';

CREATE DATABASE usersdb;

GRANT ALL PRIVILEGES ON usersdb.* TO 'myuser'@'%' WITH GRANT OPTION;

use usersdb; 

CREATE TABLE users (
    username VARCHAR(80) PRIMARY KEY,
    password VARCHAR(250),
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    address VARCHAR(2500)
);
