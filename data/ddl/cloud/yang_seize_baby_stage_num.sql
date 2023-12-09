CREATE TABLE `yang_seize_baby_stage_num` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `stage_name` varchar(255) DEFAULT NULL COMMENT '名称',
  `stageimage` varchar(255) DEFAULT NULL COMMENT '图片',
  `coin_unit_price` decimal(10,3) DEFAULT NULL COMMENT '云豆单价',
  `total_need_people_num` int(7) DEFAULT NULL COMMENT '总共需要人数',
  `surplus_people_num` int(7) DEFAULT NULL COMMENT '剩余人数',
  `participate_people_num` int(7) DEFAULT NULL COMMENT '当前参与人数',
  `least_num` int(11) DEFAULT NULL COMMENT '最少参与的云豆',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime DEFAULT NULL COMMENT '生成时间',
  `end_at` datetime DEFAULT NULL COMMENT '结束时间',
  `stage_num_type` tinyint(1) DEFAULT NULL COMMENT '类型:1=云豆夺宝,2=实物夺宝',
  `goods_id` int(11) DEFAULT NULL COMMENT '关联的商品id',
  `winning_num` int(11) DEFAULT NULL COMMENT '中奖号码',
  `stage_status` tinyint(1) DEFAULT NULL COMMENT '当前期数的状态:1=正常,2=结束,3=超时结束',
  `member_id` int(11) DEFAULT NULL COMMENT '用户id-- 发布的用户',
  `winning_user_id` int(11) DEFAULT NULL COMMENT '中奖用户的id',
  `appoint_winning_user` int(11) DEFAULT '0' COMMENT '指定的中奖用户',
  PRIMARY KEY (`id`),
  KEY `stage_status` (`stage_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='云豆夺宝的活动表';