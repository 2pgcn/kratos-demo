CREATE DATABASE IF NOT EXISTS article DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
use article;
create table article
(
    id          bigint auto_increment primary key,
    uid         bigint      not null,
    title       varchar(32)       not null comment '文章标题',
    content      text              null comment '文章内容',
    create_time bigint               not null,
    update_time bigint               not null,
    delete_time bigint               default 0 null,
    index idx_uid(uid),
    index idx_title(title(10))
) engine InnoDb;

