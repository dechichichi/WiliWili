CREATE TABLE `comments` (
    `comment_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `comment_type` TINYINT NOT NULL COMMENT '评论类型(1: 评论视频, 2: 评论评论)',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `be_comment_id` BIGINT NOT NULL COMMENT '被评论对象的ID(视频ID或评论ID)',
    `comment` TEXT NOT NULL COMMENT '评论内容',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
    `likes_count`INT NOT NULL DEFAULT 0 COMMENT '点赞数',
    PRIMARY KEY (`comment_id`),
    INDEX `idx_user_id` (`user_id`),  -- 为用户ID添加索引，便于查询
    INDEX `idx_be_comment_id` (`be_comment_id`),  -- 为被评论对象ID添加索引
    INDEX `idx_comment_type` (`comment_type`)  -- 为评论类型添加索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;