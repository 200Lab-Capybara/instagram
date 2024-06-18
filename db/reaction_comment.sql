CREATE TABLE `reaction_comments` (
     `user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `comment_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `created_at` datetime(6) NOT NULL,
     `updated_at` datetime(6) DEFAULT NULL,
     PRIMARY KEY (`comment_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci