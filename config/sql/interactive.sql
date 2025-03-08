CREATE TABLE `likes`(
    `user_id`   BIGINT NOT NULL PRIMARY KEY COMMENT '用户ID',
    `post_id`   BIGINT NOT NULL PRIMARY KEY COMMENT '帖子ID',
)
CREATE TABLE `comments`(
    `user_id`   BIGINT NOT NULL PRIMARY KEY COMMENT '用户ID',
    `post_id`   BIGINT NOT NULL PRIMARY KEY COMMENT '帖子ID',
    `comment`   TEXT NOT NULL COMMENT '评论内容'
)

