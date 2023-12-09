CREATE TABLE `yang_currency_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int(10) unsigned DEFAULT NULL COMMENT '用户id',
  `currency_id` tinyint(3) DEFAULT NULL COMMENT '资产id',
  `num` decimal(50,15) unsigned DEFAULT '0.000000000000000' COMMENT '可用数量',
  `freeze_num` decimal(50,15) unsigned DEFAULT '0.000000000000000' COMMENT '冻结数量',
  `updated_at` int(10) unsigned DEFAULT '0' COMMENT '更新时间',
  `lottery_num` decimal(50,15) NOT NULL DEFAULT '0.000000000000000' COMMENT '中奖获得的数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `member_id` (`member_id`,`currency_id`),
  KEY `member_id_2` (`member_id`),
  KEY `currency_id` (`currency_id`)
) ENGINE=InnoDB AUTO_INCREMENT=998985 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户资产表';