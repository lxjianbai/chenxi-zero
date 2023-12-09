CREATE TABLE `yang_seize_baby_winning` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `member_id` int(11) DEFAULT NULL COMMENT '用户id',
  `stage_num_id` int(11) DEFAULT NULL COMMENT '期数id',
  `winning_num` int(11) DEFAULT NULL COMMENT '中奖号码',
  `prize_des` varchar(255) DEFAULT NULL COMMENT '奖品描述',
  `prize_status` tinyint(1) DEFAULT NULL COMMENT '奖品的状态:1=未发放,2=已发放',
  `created_at` datetime DEFAULT NULL COMMENT '生成时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `prize_num` float(11,2) DEFAULT NULL COMMENT '奖品数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `stagenumid` (`stage_num_id`),
  UNIQUE KEY `onlyone` (`member_id`,`stage_num_id`),
  KEY `memberId` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户中奖表';