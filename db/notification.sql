CREATE TABLE `notifications` (
    `id` varchar(36) NOT NULL,
    `model_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `model_name` enum('Post','Story','Comment','Follow') NOT NULL,
    `user_id` varchar(36) NOT NULL,
    `from_user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `payload` json DEFAULT NULL,
    `status` enum('read','sent','created') NOT NULL, `payload` json DEFAULT NULL,
    UNIQUE KEY `id` (`id`)
)