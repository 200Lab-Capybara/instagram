create table image_detail
(
    model_name varchar(255) not null,
    image_id   varchar(36)  not null,
    model_id   varchar(36)  not null,
    primary key (model_id, image_id)
);

