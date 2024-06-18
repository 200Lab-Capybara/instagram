create table images
(
    id           varchar(36)             not null
        primary key,
    image_url    varchar(255)            not null,
    size         int                     not null,
    width        int                     not null,
    height       int                     not null,
    status       enum ('used', 'unused') not null,
    create_at    datetime(6)             not null,
    updated_at   datetime(6)             null,
    storage_name varchar(255)            not null,
    user_id      varchar(36)             not null
);

