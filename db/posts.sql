CREATE TABLE `posts` (
     `id` varchar(36) NOT NULL DEFAULT '',
     `user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `content` varchar(1000) DEFAULT NULL,
     `like_count` int NOT NULL DEFAULT '0',
     `comment_count` int NOT NULL DEFAULT '0',
     `status` enum('active','deleted') NOT NULL DEFAULT 'active',
     `used_hashtag` tinyint(1) NOT NULL DEFAULT '0',
     `created_at` datetime(6) DEFAULT NULL,
     `updated_at` datetime(6) DEFAULT NULL,
     PRIMARY KEY (`id`),
     KEY `posts_user_id_idx` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;