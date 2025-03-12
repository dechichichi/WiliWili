CREATE TABLE `video` (
                        `id` BIGINT NOT NULL PRIMARY KEY COMMENT '视频ID',
                        `title` VARCHAR(255) NOT NULL COMMENT '视频标题',
                        `url` VARCHAR(255) NOT NULL COMMENT '视频地址',
                        `duration` INT NOT NULL COMMENT '视频时长',
                        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='视频表';
