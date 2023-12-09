CREATE TABLE `yang_seize_baby_activity_people_num` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `member_id` int(11) DEFAULT NULL COMMENT '用户id',
  `activity_id` int(11) DEFAULT NULL COMMENT '活动id',
  `frequency_num` int(11) DEFAULT NULL COMMENT '参与活动的人次数',
  `created_at` datetime DEFAULT NULL COMMENT '开始参与活动的时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后参与活动的时间',
  `is_winning` tinyint(1) DEFAULT NULL COMMENT '是否中奖:1=中奖,3=未中奖,2=开奖进行中',
  PRIMARY KEY (`id`),
  KEY `member_id` (`member_id`),
  KEY `activityId` (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='活动参与人次表';