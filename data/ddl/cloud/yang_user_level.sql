CREATE TABLE `yang_user_level` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户id',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级(列出所有的父级id)',
  `level` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '相对于父级的等级',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_pid` (`user_id`,`pid`),
  KEY `idx_pid_userid` (`pid`,`user_id`),
  KEY `level` (`level`),
  KEY `user_level` (`user_id`,`level`),
  KEY `pid` (`pid`,`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='邀请人等级关系表';