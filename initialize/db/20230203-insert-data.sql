-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO `tb_sys_user` (`id`,
                           `created_at`,
                           `updated_at`,
                           `deleted_at`,
                           `uuid`,
                           `mobile`,
                           `password`,
                           `avatar`,
                           `name`,
                           `status`,
                           `role_id`,
                           `last_login`,
                           `locked`,
                           `lock_expire`,
                           `pass_wrong`)
VALUES (1,
        '2023-02-03 11:01:40.848',
        '2023-02-03 11:01:40.848',
        NULL,
        'f6784ed2-6d7a-4d1b-ae97-8b909100eab6',
        '18111111111',
        '$2a$10$vFRuEvEh4oNOpFvzau.BH.sNp8ffm6PIYeL2NzH4Ng853g4VIeUsa',
        'http://pic.imeitou.com/uploads/allimg/230201/7-230201162K0.jpg',
        '无敌小老鼠',
        0,
        0,
        NULL,
        0,
        0,
        0);


INSERT INTO `tb_sys_role` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `keyword`, `desc`, `status`, `sort`)
VALUES (1, '2023-02-03 11:01:40', '2023-02-03 11:01:40', NULL, '超级管理员', 'super', '拥有所有管理权限的超级角色', 0,
        0),
       (2, '2023-02-03 11:01:40', '2023-02-03 11:01:40', NULL, '测试用户', 'guest', '拥有访问部分资源的测试角色', 0, 1);