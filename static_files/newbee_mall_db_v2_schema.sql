
# 设置当前会话的字符集为 utf8mb4
SET NAMES utf8mb4;

# 关闭当前会话中的外键检查。当执行一些需要修改数据库结构的操作时（比如数据迁移或数据导入），关闭外键检查可以避免因外键约束而导致的问题。
# 这样可以在处理数据时更加灵活，但需要谨慎使用，以确保数据的一致性。
SET FOREIGN_KEY_CHECKS = 0;

# 创建数据库

DROP DATABASE IF EXISTS newbee_mall_db_v2;

# utf8mb4_general_ci 表示使用 UTF-8 编码，支持多语言和特殊字符，同时在比较和排序文本时不区分大小写。这通常是一个合适的默认选项，适用于大多数情况下的文本存储和处理需求。
CREATE DATABASE IF NOT EXISTS newbee_mall_db_v2 DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;;

# Dump of table tb_newbee_mall_admin_user
# ------------------------------------------------------------
DROP TABLE IF EXISTS `tb_newbee_mall_admin_user`;
CREATE TABLE IF NOT EXISTS `tb_newbee_mall_admin_user`(
    `admin_user_id` bigint(20) AUTO_INCREMENT COMMENT '管理员id',
    `login_user_name` varchar(50) NOT NULL COMMENT '管理员登录名称',
    `login_password` varchar(50) NOT NULL COMMENT '管理员登录密码',
    `nick_name` varchar(50) NOT NULL COMMENT '管理员显示昵称',
    `locked` tinyint(4) DEFAULT '0' COMMENT '是否锁定，0未锁定 1已锁定无法登陆',
    PRIMARY KEY (admin_user_id) USING BTREE #索引类型，是一种常见的用于排序和搜索的树状结构。使用 BTREE 索引可以加速查询操作。
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC; # DYNAMIC：使用动态行格式。动态行格式可以根据数据的大小和特性自动选择行的存储方式，以优化存储空间。

# 表示以写入（修改）模式锁定表。这意味着在锁定期间，其他会话无法对表进行写入（更新、插入、删除）操作。读操作一般是允许的，但写操作将被阻塞。
LOCK TABLES `tb_newbee_mall_admin_user` WRITE;

# /*!40000: 这是 MySQL 中的条件注释，用于指定执行这个 SQL 语句的最低版本。在这个例子中，条件注释 40000 表示 MySQL 版本需大于等于 4.0.0 才会执行这个语句
#  DISABLE KEYS：表示要禁用表上的索引更新。当执行大量数据插入操作时，禁用索引更新可以显著提高性能。
/*!40000 ALTER TABLE `tb_newbee_mall_admin_user` DISABLE KEYS */;

INSERT INTO `tb_newbee_mall_admin_user` (`admin_user_id`, `login_user_name`, `login_password`, `nick_name`, `locked`)
VALUES
    (1,'admin','e10adc3949ba59abbe56e057f20f883e','十三',0),
    (2,'newbee-admin1','e10adc3949ba59abbe56e057f20f883e','新蜂01',0),
    (3,'newbee-admin2','e10adc3949ba59abbe56e057f20f883e','新蜂02',0);


# 当数据操作完成后，重新启用索引更新。
/*!40000 ALTER TABLE `tb_newbee_mall_admin_user` ENABLE KEYS */;

UNLOCK TABLES;


# Dump of table tb_newbee_mall_carousel 轮播图
# ------------------------------------------------------------
DROP TABLE IF EXISTS `tb_newbee_mall_carousel`;

CREATE TABLE IF NOT EXISTS `tb_newbee_mall_carousel`(
    `carousel_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '首页轮播图id',
    `carousel_url` varchar(100) NOT NULL DEFAULT '' COMMENT '轮播图',
    `redirect_url` varchar(100) NOT NULL DEFAULT '''##''' COMMENT '点击后的跳转地址(默认不跳转)',

) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;



