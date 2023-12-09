CREATE TABLE `yang_seize_baby_member_stage_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `member_id` int(11) DEFAULT NULL COMMENT '用户id',
  `stage_num_id` int(11) DEFAULT NULL COMMENT '期数id',
  `winning_num` int(11) DEFAULT NULL COMMENT '开奖起始号码',
  `end_winning_num` int(11) DEFAULT NULL COMMENT '开奖结束号码',
  `participate_num` int(5) DEFAULT NULL COMMENT '参与的云豆数量',
  `status` tinyint(1) DEFAULT NULL COMMENT '当前订单的状态:1=正常,2=结束,3=退款,4=超时退款',
  `created_at` datetime DEFAULT NULL COMMENT '生成时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `winning_probability` int(3) DEFAULT NULL COMMENT '中奖概率%',
  `is_winning` tinyint(1) DEFAULT NULL COMMENT '是否中奖:1=中奖,3=未中奖,2=开奖进行中',
  `open_winning_at` datetime DEFAULT NULL COMMENT '开奖时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `numid` (`winning_num`,`stage_num_id`),
  UNIQUE KEY `numstartend` (`stage_num_id`,`winning_num`,`end_winning_num`),
  KEY `stagenumid` (`stage_num_id`),
  KEY `memberid` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='夺宝人数表';