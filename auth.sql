CREATE DATABASE IF NOT EXISTS auth DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
create table user
(
    id          bigint auto_increment
        primary key,
    username    varchar(32)       not null
        unique,
    password    varchar(32)       not null,
    avatar      text              null comment '用户头像',
    status      tinyint default 1 not null comment '1 正常,2 封禁',
    create_time int               not null,
    delete_time int     default 0 null
);

