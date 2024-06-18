create table stories
(
    id            varchar(36)   not null
        primary key,
    user_id       varchar(36)   not null,
    content_story varchar(255)  null,
    react_count   int default 0 not null,
    expires_time  int           null,
    created_at    datetime(6)   null,
    updated_at    datetime(6)   null
);

