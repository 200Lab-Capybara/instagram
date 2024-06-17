create table stories
(
    id               varchar(36)                         not null
        primary key,
    reaction_user_id varchar(36)                         not null,
    text_story       varchar(255)                        null,
    user_id          varchar(36)                         not null,
    create_at        timestamp default CURRENT_TIMESTAMP null,
    expires_at       timestamp                           null,
    like_count       int       default 0                 not null
);

