CREATE TABLE `user_followers` (
  `user_id` varchar(36) NOT NULL,
  `following` varchar(36) NOT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`following`,`user_id`),
  KEY `user_followers_user_id_idx` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;