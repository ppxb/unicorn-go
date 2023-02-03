-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `tb_sys_role` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `keyword`, `desc`, `status`, `sort`)
VALUES (1, '2023-02-03 11:01:40', '2023-02-03 11:01:40', NULL, '超级管理员', 'super', '拥有所有管理权限的超级角色', 0,
        0),
       (2,
        '2023-02-03 11:01:40',
        '2023-02-03 11:01:40',
        NULL,
        '测试用户',
        'guest',
        '拥有访问部分资源的测试角色',
        0,
        1);