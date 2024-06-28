CREATE TABLE `notifications` (
     `id` varchar(36) NOT NULL,
     `model_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `model_name` enum('Post','Story','Comment','Follow') NOT NULL,
     `user_id` varchar(36) NOT NULL,
     `from_user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `created_at` datetime(6) NOT NULL,
     `payload` json DEFAULT NULL,
     `status` enum('read','sent','created') NOT NULL,
     `updated_at` datetime(6) DEFAULT NULL,
     PRIMARY KEY (`id`),
     KEY `user_id` (`user_id`) USING BTREE,
     KEY `notifications_created_at_idx` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;