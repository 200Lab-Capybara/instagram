CREATE TABLE `stories`
(
    `id`            varchar(36) NOT NULL,
    `user_id`       varchar(36) NOT NULL,
    `content_story` varchar(255)         DEFAULT NULL,
    `react_count`   int         NOT NULL DEFAULT '0',
    `expires_time`  int                  DEFAULT NULL,
    `created_at`    datetime(6) DEFAULT NULL,
    `updated_at`    datetime(6) DEFAULT NULL,
    `image_id`      varchar(36) NOT NULL,
    `is_active`     tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `stories_user_id_idx` (`user_id`) USING BTREE,
    UNIQUE KEY `stories_created_at_idx` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;