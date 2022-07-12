CREATE DATABASE IF NOT EXISTS ps-user;
USE ps-user;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createIn timestamp default current_timestamp()
) ENGINE=INNODB;



