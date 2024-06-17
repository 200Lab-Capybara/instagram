create table images
(
    id           varchar(36)                         not null
        primary key,
    file_name    varchar(255)                        not null,
    size         int                                 not null,
    width        int                                 not null,
    height       int                                 not null,
    status       enum ('use', 'unUse')               not null,
    create_at    timestamp default CURRENT_TIMESTAMP not null,
    updated_at   timestamp                           null,
    storage_name varchar(255)                        not null,
    user_id      varchar(36)                         not null
);

