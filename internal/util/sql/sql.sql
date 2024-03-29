-- POSTGRES

CREATE TABLE users(
    id serial primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createIn timestamp default current_timestamp
)


-- MySQL
CREATE DATABASE IF NOT EXISTS psusers;
USE psusers;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createIn timestamp default current_timestamp()
) ENGINE=INNODB;
