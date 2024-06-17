CREATE TABLE `profiles` (
    `user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `date_of_birth` date DEFAULT NULL,
    `sex` enum('Male','Female','Other') NOT NULL DEFAULT 'Other',
    `avatar` json DEFAULT NULL,
    `count_following` int NOT NULL DEFAULT '0',
    `count_followers` int NOT NULL DEFAULT '0',
    `count_posts` int NOT NULL DEFAULT '0',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci