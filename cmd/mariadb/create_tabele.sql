--  Drop table

DROP TABLE IF EXISTS magazine.m_inventory;

CREATE TABLE `m_inventory` (
  `id` bigint(20) unsigned NOT NULL,
  `rel_type` bigint(20) unsigned NOT NULL COMMENT '实体类型',
  `rel_id` bigint(20) unsigned NOT NULL COMMENT '实体',
  `qty` int(10) unsigned NOT NULL DEFAULT 0,
  `rest_qty` int(10) unsigned NOT NULL DEFAULT 0,
  `version` int(10) unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存总表';


DROP TABLE IF EXISTS magazine.m_inventory_incr;
CREATE TABLE `m_inventory_incr` (
  `id` bigint(20) unsigned NOT NULL,
  `change_qty` int(10) unsigned NOT NULL DEFAULT 0,
  `snopshot` varchar(256) NOT NULL DEFAULT '库存快照',
  `tran_id` bigint(20) unsigned NOT NULL COMMENT '库存流水明细记录ID',
  `incr_type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '入库类型(1 库存初始化, 2 库存返换)',
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存入库表';


DROP TABLE IF EXISTS  magazine.m_inventory_decr;
CREATE TABLE if NOT EXISTS `m_inventory_decr` (
  `id` bigint(20) unsigned NOT NULL,
  `change_qty` int(10) unsigned NOT NULL DEFAULT 0,
  `snopshot` varchar(256) NOT NULL DEFAULT '库存快照',
  `tran_id` bigint(20) unsigned NOT NULL COMMENT '库存流水明细记录ID',
  `incr_type` tinyint(4) NOT NULL DEFAULT 1 COMMENT "出库类型(1 普通库存扣减)",
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存出库表';

DROP TABLE IF EXISTS  magazine.m_inventory_tran;
CREATE TABLE if NOT EXISTS  `m_inventory_tran` (
  `id` bigint(20) unsigned NOT NULL,
  `changet_type` tinyint(4) unsigned  not null default 0 COMMENT "入库类型(1 出库, 2 入库)",
  `change_qty` int(11) unsigned NOT NULL DEFAULT '0',
  `tran_id` bigint(20) unsigned not null COMMENT "库存流水明细记录ID",
  `tran_status`  tinyint(4) unsigned not null default 1 COMMENT "明细状态 (1 处理中，2 已完成，3 处理失败)",
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  `remark`  varchar(256) not null default "备注" ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存明细申请记录'
