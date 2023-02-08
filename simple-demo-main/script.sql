create table comment
(
    id          bigint                              not null
        primary key,
    user_id     bigint                              null,
    video_id    bigint                              null,
    content     text                                null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp                           null
)
    comment '评论表';

create table osskey
(
    `key`  varchar(50) null,
    secret varchar(50) null
);

create table user
(
    id             bigint        not null
        primary key,
    name           longtext      null,
    follow_count   bigint        null,
    follower_count bigint        null,
    is_follow      tinyint(1)    null,
    user_id        varchar(40)   null,
    password       varchar(40)   null,
    slat           int default 0 null
);

create table user_account
(
    id       int auto_increment
        primary key,
    userid   int         null,
    username varchar(50) null,
    password varchar(50) null,
    salt     int         null
);

create table video
(
    id             bigint     not null
        primary key,
    author_id      bigint     null,
    play_url       longtext   null,
    cover_url      longtext   null,
    favorite_count bigint     null,
    comment_count  bigint     null,
    is_favorite    tinyint(1) null
);


