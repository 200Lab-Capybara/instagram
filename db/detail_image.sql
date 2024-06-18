create table image_detail
(
    model_name varchar(255) not null,
    image_id   varchar(36)  not null,
    model_id   varchar(36)  not null,
    create_at  datetime(6)  not null,
    updated_at datetime(6)  null,
    primary key (image_id, model_id)
);

