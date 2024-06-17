create table story_likes
(
    id         varchar(36)                          not null
        primary key,
    user_id    varchar(36)                          not null,
    story_id   varchar(36)                          not null,
    liked      tinyint(1) default 0                 not null,
    created_at timestamp  default CURRENT_TIMESTAMP not null
);

