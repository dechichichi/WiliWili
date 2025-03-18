CREATE TABLE `video` (
    `video_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `video_name` VARCHAR(255) NOT NULL COMMENT '视频名称',
    `video_url` VARCHAR(500) NOT NULL COMMENT '视频链接',
    `video_duration` VARCHAR(50) NOT NULL COMMENT '视频时长',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;