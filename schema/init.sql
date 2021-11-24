create table books
(
    id          serial       not null unique,
    name        varchar(100) not null,
    title       varchar(100) not null,
    isbn        varchar(100) not null unique,
    description text         not null
);