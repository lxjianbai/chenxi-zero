CREATE TABLE `yang_active_level_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户主键',
  `source_uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '变更来源用户主键',
  `event` varchar(50) NOT NULL DEFAULT '' COMMENT '事件标识',
  `rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '参与计算的比例',
  `result_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '运算后的比例',
  `link_id` varchar(255) NOT NULL DEFAULT '' COMMENT '关联主键',
  `pm` tinyint(1) NOT NULL DEFAULT '1' COMMENT '增减类别。0：减少，1：增加',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注内容',
  `ymd` char(8) NOT NULL DEFAULT '' COMMENT '年月日',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='活跃度记录表日志表';