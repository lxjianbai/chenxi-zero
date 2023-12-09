CREATE TABLE `yang_seize_baby_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `goods_name` varchar(255) DEFAULT NULL COMMENT '商品的名字',
  `goods_price` decimal(7,2) DEFAULT NULL COMMENT '商品的价格',
  `goodsimage` varchar(150) DEFAULT NULL COMMENT '商品图片',
  `created_at` datetime DEFAULT NULL COMMENT '商品添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '商品更新时间',
  `goods_status` tinyint(1) DEFAULT NULL COMMENT '商品的状态:1=上架,2=下架',
  `goods_type` tinyint(4) DEFAULT NULL COMMENT '商品的类型:1=云豆夺宝,2=实物夺宝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='云豆夺宝中的商品表';