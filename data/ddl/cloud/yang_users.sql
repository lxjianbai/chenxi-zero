CREATE TABLE `yang_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'user id.',
  `name` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `email` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'user email',
  `phone` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'user phone member.',
  `avatar` varchar(220) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
  `location` varchar(180) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户位置',
  `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'password.',
  `bio` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户简介',
  `sex` tinyint(4) DEFAULT '0' COMMENT '用户性别，1男  2女  0保密',
  `create_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  `pay_password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付密码',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `ip` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP',
  `version` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '版本',
  `robot_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态机器人，1',
  `vip_grade` int(10) DEFAULT '0' COMMENT '用户VIP等级',
  `is_delete` tinyint(1) unsigned zerofill DEFAULT '0' COMMENT '是否删除：0=正常用户，1=删除用户',
  `disable_login_time` int(11) DEFAULT '0' COMMENT '禁用登录到期时间戳',
  `lock_mark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '禁用备注',
  `is_certification` tinyint(2) DEFAULT '0' COMMENT '实名认证：0=未实名，1实名',
  `share_certification` tinyint(1) DEFAULT '0' COMMENT '分享认证：0=未认证，1=认证',
  `senio_certification` tinyint(1) DEFAULT '0' COMMENT '高级会员认证：0=未认证，1=认证',
  `top_certification` tinyint(1) DEFAULT '0' COMMENT '顶级认证：0=未认证，1=认证',
  `operator_certification` tinyint(1) DEFAULT '0' COMMENT '运营商认证，1.县级  2.市级 3.省级',
  `disable_otc_time` int(11) DEFAULT '0' COMMENT '禁用otc到期时间戳',
  `disable_otc_mark` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '禁用OTC备注',
  `is_upgrade_pwd` tinyint(2) unsigned DEFAULT '1' COMMENT '是否更新密码',
  `is_active` tinyint(1) DEFAULT '0' COMMENT '是否为有效会员,0=无效，1=有效',
  `user_type` tinyint(4) DEFAULT '0' COMMENT '用户类型，0.正常注册  1.快速注册',
  `manual_update` tinyint(2) unsigned DEFAULT '2' COMMENT '是否手动更新等级:1是，2否',
  `user_mark` mediumint(5) unsigned DEFAULT '0' COMMENT '身份标识：0普通用户1联创2市场部，关联mark表',
  `openid` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'openid',
  `unionid` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'unionid',
  `alipay_userid` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付宝用户ID',
  `delete_mark` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '禁用备注',
  `studio_id` mediumint(5) unsigned DEFAULT '0' COMMENT '工作室列表id',
  `studio_edit_num` int(11) DEFAULT '1' COMMENT '服务中心工作室修改次数',
  `dym_openid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '通过本表的id使用md5加密生成的参数',
  `is_clear` tinyint(4) NOT NULL DEFAULT '2' COMMENT '是否清理过，1清理过0没清理',
  `certification_time` int(11) DEFAULT '0' COMMENT '实名认证时间',
  `clear_mark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '15天未登录清理备注',
  `is_show` int(11) DEFAULT '1' COMMENT '是否显示：1：是，0：否',
  `good_number` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '靓号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_name_unique` (`name`),
  UNIQUE KEY `users_phone_unique` (`phone`,`name`,`create_time`),
  UNIQUE KEY `users_email_unique` (`email`),
  UNIQUE KEY `good_number_unique` (`good_number`),
  UNIQUE KEY `users_phone` (`phone`),
  KEY `share_cert` (`id`,`is_certification`,`is_active`),
  KEY `is_active` (`is_active`,`vip_grade`),
  KEY `last_time` (`last_time`),
  KEY `vip_grade` (`vip_grade`),
  KEY `studio_id` (`studio_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户';