CREATE TABLE `likes` (
                         `user_id` BIGINT NOT NULL COMMENT '用户ID',
                         `post_id` BIGINT NOT NULL COMMENT '帖子ID',
                         PRIMARY KEY (`user_id`, `post_id`)  -- 复合主键
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;






CREATE TABLE `comments` (
                            `user_id` BIGINT NOT NULL COMMENT '用户ID',
                            `post_id` BIGINT NOT NULL COMMENT '帖子ID',
                            `comment` TEXT NOT NULL COMMENT '评论内容',
                            PRIMARY KEY (`user_id`, `post_id`)  -- 复合主键
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


