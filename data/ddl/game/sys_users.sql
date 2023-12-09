CREATE TABLE `sys_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(20) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '名称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 1=正常 2:锁定',
  `role_ids` varchar(10) NOT NULL DEFAULT '' COMMENT '角色id',
  `login_at` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_account_status` (`account`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;