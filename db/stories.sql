CREATE TABLE `stories`
(
    `id`            varchar(36) NOT NULL,
    `user_id`       varchar(36) NOT NULL,
    `image_id`      varchar(36) NOT NULL,
    `content_story` varchar(255)         DEFAULT NULL,
    `react_count`   int         NOT NULL DEFAULT '0',
    `expires_time`  int                  DEFAULT NULL,
    `created_at`    datetime(6) DEFAULT NULL,
    `updated_at`    datetime(6) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;