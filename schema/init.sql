create table books
(
    id          serial       not null unique,
    name        varchar(100) not null,
    title       varchar(100) not null,
    isbn        varchar(100) not null unique,
    description text         not null
);

create table users
(
    id            serial       not null unique,
    name          varchar(100) not null,
    email         varchar(100) not null unique,
    password_hash varchar(100) not null
);