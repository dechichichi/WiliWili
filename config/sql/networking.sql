CREATE TABLE `networking` (
    `user_id`   BIGINT NOT NULL PRIMARY KEY COMMENT '用户ID',
    `followers` JSON NOT NULL COMMENT '粉丝列表',
    `following` JSON NOT NULL COMMENT '关注列表',
    `friends`   JSON NOT NULL COMMENT '好友列表'        
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户关系表';