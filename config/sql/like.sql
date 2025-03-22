CREATE TABLE `video_like` (
                         `user_id` BIGINT NOT NULL COMMENT '用户ID',
                         `video_id` BIGINT NOT NULL COMMENT '视频ID',
                        `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         PRIMARY KEY (`user_id`, `video_id`)  -- 复合主键
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `comment_like` (
                         `user_id` BIGINT NOT NULL COMMENT '用户ID',
                         `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `comment_id` BIGINT NOT NULL COMMENT '评论ID',
                         PRIMARY KEY (`user_id`, `comment_id`)  -- 复合主键
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
