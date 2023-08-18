CREATE DATABASE IF NOT EXISTS auth DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
use auth;
create table user
(
    id          bigint auto_increment primary key,
    username    varchar(32)       not null unique key ,
    password    varchar(32)       not null,
    avatar      text              null comment '用户头像',
    email       varchar(128)      not null comment 'email',
    status      tinyint default 1 not null comment '1 正常,2 封禁',
    create_time bigint               not null,
    delete_time bigint               default 0 null,
    index idx_email(email(10))
) engine InnoDb;

