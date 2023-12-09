CREATE TABLE `yang_seize_baby_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `key` varchar(255) DEFAULT NULL COMMENT 'key',
  `value` varchar(255) DEFAULT NULL COMMENT 'value',
  `describe` varchar(300) DEFAULT NULL COMMENT '描述',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态:1=启用,0禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='多人夺宝配置表';