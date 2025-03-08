CREATE TABLE 'users' (
      `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `username` VARCHAR(30) NOT NULL COMMENT '用户名最多 10 个中文字符或等长英文字符',
    `password`VARCHAR(255) NOT NULL COMMENT '数字+字母组合，总长度上限 16',
    `email` VARCHAR(50) NOT NULL COMMENT '邮箱',
    `phone` VARCHAR(11) NOT NULL COMMENT '手机号',
    `signature` VARCHAR(255) NOT NULL COMMENT '签名,最多 200 个字符',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `gender` ENUM('male', 'female', 'other') NOT NULL COMMENT '性别',
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;