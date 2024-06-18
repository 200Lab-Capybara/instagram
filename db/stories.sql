create table stories
(
    id            varchar(36)                         not null
        primary key,
    content_story varchar(255)                        null,
    user_id       varchar(36)                         not null,
    created_at    timestamp default CURRENT_TIMESTAMP null,
    expires_time  timestamp                           null,
    react_count   int       default 0                 not null,
    updated_at    timestamp                           null
);

