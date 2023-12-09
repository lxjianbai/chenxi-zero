CREATE TABLE `sys_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `unique_key` varchar(50) NOT NULL DEFAULT '' COMMENT '唯一标识',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `menu_ids` varchar(100) NOT NULL DEFAULT '' COMMENT '权限集',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_key` (`unique_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';