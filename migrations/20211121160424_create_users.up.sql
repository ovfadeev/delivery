CREATE TABLE users
(
    id         serial       not null primary key,
    created_at timestamp    not null default current_timestamp,
    login      varchar(50)  not null,
    key        varchar(100) not null
)