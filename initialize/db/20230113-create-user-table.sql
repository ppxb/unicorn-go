-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE
    IF
    NOT EXISTS `tb_sys_user`
(
    `id`
    BIGINT
    UNSIGNED
    NOT
    NULL
    AUTO_INCREMENT
    COMMENT
    "auto increment id",
    `created_at`
    DATETIME
(
    3
) DEFAULT NULL COMMENT "创建时间",
    `updated_at` DATETIME
(
    3
) DEFAULT NULL COMMENT "更新时间",
    `deleted_at` DATETIME
(
    3
) DEFAULT NULL COMMENT "软删除时间",
    `mobile` VARCHAR
(
    20
) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT "用户手机",
    `password` VARCHAR
(
    255
) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT "用户密码",
    PRIMARY KEY
(
    `id`
),
    UNIQUE KEY `idx_mobile`
(
    `mobile`
),
    KEY `idx_deleted_at`
(
    `deleted_at`
)
    ) ENGINE = INNODB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;