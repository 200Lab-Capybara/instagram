create table react_story
(
    user_id    varchar(36) not null,
    story_id   varchar(36) not null,
    created_at datetime(6) not null,
    updated_at datetime(6) null,
    primary key (story_id, user_id)
);

