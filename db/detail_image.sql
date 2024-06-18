create table image_detail
(
    model_name varchar(255)                        not null,
    image_id   varchar(36)                         not null,
    model_id   varchar(36)                         not null,
    create_at  timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp                           null,
    primary key (image_id, model_id)
);

