-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE
    IF
    NOT EXISTS `tb_sys_user`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自动递增ID',
    `created_at`  DATETIME(3)                             DEFAULT NULL COMMENT '创建时间',
    `updated_at`  DATETIME(3)                             DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  DATETIME(3)                             DEFAULT NULL COMMENT '软删除时间',
    `uuid`        VARCHAR(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'uuid',
    `mobile`      VARCHAR(20) COLLATE utf8mb4_general_ci COMMENT '用户手机号码',
    `password`    LONGTEXT COLLATE utf8mb4_general_ci COMMENT '用户密码',
    `avatar`      LONGTEXT COLLATE utf8mb4_general_ci COMMENT '头像地址',
    `name`        LONGTEXT COLLATE utf8mb4_general_ci COMMENT '用户姓名',
    `status`      TINYINT(1)                              DEFAULT '0' COMMENT '账号状态(0: 未禁用, 1: 已禁用)',
    `role_id`     BIGINT UNSIGNED                         DEFAULT NULL COMMENT '角色ID',
    `last_login`  DATETIME(3)                             DEFAULT NULL COMMENT '上次登录时间',
    `locked`      TINYINT(1)                              DEFAULT '0' COMMENT '锁定状态(0: 未锁定, 1: 已锁定)',
    `lock_expire` BIGINT                                  DEFAULT NULL COMMENT '锁定剩余时间',
    `pass_wrong`  BIGINT                                  DEFAULT NULL COMMENT '密码错误次数',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uuid` (`uuid`),
    KEY `idx_mobile` (`mobile`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE = INNODB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT '用户表';

CREATE TABLE
    IF
    NOT EXISTS `tb_sys_role`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自动递增ID',
    `created_at` DATETIME(3)                             DEFAULT NULL COMMENT '创建时间',
    `updated_at` DATETIME(3)                             DEFAULT NULL COMMENT '更新时间',
    `deleted_at` DATETIME(3)                             DEFAULT NULL COMMENT '软删除时间',
    `name`       LONGTEXT COLLATE utf8mb4_general_ci COMMENT '名称',
    `keyword`    VARCHAR(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关键字(唯一)',
    `desc`       LONGTEXT COLLATE utf8mb4_general_ci COMMENT '描述',
    `status`     TINYINT(1)                              DEFAULT '0' COMMENT '状态(0: 未禁用, 1: 已禁用)',
    `sort`       BIGINT UNSIGNED                         DEFAULT '1' COMMENT '排序(值越小权限越大，应>=0。0：超级管理员)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_keyword` (`keyword`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE = INNODB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT '角色表';