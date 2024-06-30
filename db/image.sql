CREATE TABLE `images`
(
    `id`           varchar(36)  NOT NULL,
    `image_url`    varchar(255) NOT NULL,
    `size`         int          NOT NULL,
    `width`        int          NOT NULL,
    `height`       int          NOT NULL,
    `status`       enum('used','unused') NOT NULL,
    `create_at`    datetime(6) NOT NULL,
    `updated_at`   datetime(6) DEFAULT NULL,
    `storage_name` varchar(255) NOT NULL,
    `user_id`      varchar(36)  NOT NULL,
    `extension`    varchar(36)  NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
