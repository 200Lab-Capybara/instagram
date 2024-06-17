CREATE TABLE hashtags (
    `id` varchar(36) NOT NULL,
    `tag` VARCHAR(255) NOT NULL UNIQUE,
    `post_id` varchar(36) NOT NULL,
    `created_at` timestamp(6) NULL DEFAULT NULL,
    `updated_at` timestamp(6) NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
);
