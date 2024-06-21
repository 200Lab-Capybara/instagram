CREATE TABLE `users` (
     `id` varchar(36) NOT NULL,
     `email` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `password` varchar(255) NOT NULL,
     `salt` varchar(50) NOT NULL,
     `first_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `last_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `status` enum('active','inactive','banned', 'deleted') DEFAULT 'active',
     `role` enum('user','moderator','admin') DEFAULT 'user',
     `created_at` datetime(6) NULL DEFAULT NULL,
     `updated_at` datetime(6) NULL DEFAULT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;