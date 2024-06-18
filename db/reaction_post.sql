CREATE TABLE `reaction_posts` (
    `post_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `created_at` datetime(6) NOT NULL,
    PRIMARY KEY (`post_id`,`user_id`)
)