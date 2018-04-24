# 创建数据库
CREATE DATABASE `link` /*!40100 DEFAULT CHARACTER SET utf8 */;

# 创建 category 表
CREATE TABLE `link`.`category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` int(11) UNIQUE,
  `category_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# 创建 link 表
CREATE TABLE `link`.`link` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(200) NOT NULL,
  `category_id` int(11) NOT NULL DEFAULT '1',
  `tag` varchar(200),
  PRIMARY KEY (`id`),
  CONSTRAINT `category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
#------ 说明 ------
# 1. link 表的 category_id 默认值为 1, 即是 "未分类"
# 2. tag 使用 ","(英文逗号) 分割

#------ 初始化 ------
# 初始化 "未分类" 类别
INSERT INTO `link`.`category` (category_id,category_name) VALUES (1, "未分类");
