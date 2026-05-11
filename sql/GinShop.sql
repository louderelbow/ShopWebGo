-- MySQL dump 10.13  Distrib 8.0.45, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: ginshop
-- ------------------------------------------------------
-- Server version	8.0.45

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `access`
--

DROP TABLE IF EXISTS `access`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `access` (
  `id` int NOT NULL AUTO_INCREMENT,
  `module_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '0',
  `type` tinyint(1) DEFAULT NULL,
  `action_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `module_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `access`
--

LOCK TABLES `access` WRITE;
/*!40000 ALTER TABLE `access` DISABLE KEYS */;
INSERT INTO `access` VALUES (52,'管理员管理',1,'','',0,105,'管理员管理',0,1),(53,'角色管理',1,'','',0,100,'角色管理',0,1),(54,'管理员管理',2,'管理员列表','manager',52,100,'管理员列表',0,1),(55,'管理员管理',2,'增加管理员','manager/add',52,100,'管理员列表',0,1),(56,'管理员管理',3,'编辑管理员','manager/edit',52,100,'编辑管理员',0,1),(57,'管理员管理',3,'删除管理员','manager/delete',52,100,'删除管理员',0,1),(59,'角色管理',2,'角色列表','role',53,100,'角色列表',0,1),(60,'角色管理',2,'增加角色','role/add',53,100,'增加角色',0,1),(61,'角色管理',3,'编辑角色','role/edit',53,100,'编辑角色',0,1),(62,'角色管理',3,'删除角色','role/delete',53,100,'删除角色',0,1),(63,'权限管理',1,'','',0,100,'权限管理',0,1),(64,'权限管理',2,'权限列表','access',63,100,'',0,1),(67,'权限管理',2,'增加权限','access/add',63,100,'',0,1),(68,'轮播图管理',1,'','',0,100,'',0,1),(69,'轮播图管理',2,'轮播图列表','focus',68,101,'1111',0,1),(70,'轮播图管理',2,'增加轮播图','focus/add',68,100,'增加轮播图',0,1),(71,'轮播图管理',3,'编辑轮播图','focus/edit',68,100,'',0,1),(75,'轮播图管理',3,'删除轮播图','focus/delete',68,100,'',0,1),(76,'管理员管理',3,'执行增加管理员','manager/doAdd',52,100,'执行增加',0,1),(77,'管理员管理',3,'执行修改管理员','manager/doEdit',52,100,'执行修改',0,1),(78,'角色管理',3,'执行增加角色','role/doAdd',53,100,'执行增加',0,1),(79,'角色管理',3,'执行修改角色','role/doEdit',53,100,'执行修改',0,1),(80,'角色管理',3,'角色授权','role/auth',53,100,'',0,1),(81,'角色管理',3,'执行角色授权','role/doAuth',53,100,'执行授权',0,1),(82,'权限管理',3,'修改权限','access/edit',63,100,'执行修改',0,1),(83,'权限管理',3,'删除权限','access/delete',63,100,'',0,1),(84,'权限管理',3,'执行增加权限','access/doAdd',63,100,'',0,1),(85,'权限管理',3,'执行修改权限','access/doEdit',63,100,'执行修改\r\n',0,1),(86,'轮播图管理',3,'执行增加','focus/doAdd',68,100,'',0,1),(87,'商品管理',1,'','',0,100,'',0,1),(88,'商品管理',2,'商品分类列表','goodsCate',87,100,'',0,1),(89,'商品管理',3,'增加商品分类','goodsCate/add',87,100,'',0,1),(90,'商品管理',3,'执行增加商品分类','goodsCate/doAdd',87,100,'',0,1),(91,'商品管理',3,'修改商品分类','goodsCate/edit',87,100,'',0,1),(92,'商品管理',3,'执行修改商品分类','goodsCate/doEdit',87,100,'',0,1),(93,'商品管理',3,'删除商品分类','goodsCate/delete',87,100,'',0,1),(94,'商品管理',2,'商品类型列表','goodsType',87,100,'',0,1),(95,'商品管理',3,'增加商品类型','goodsType/add',87,100,'',0,1),(96,'商品管理',3,'编辑商品类型','goodsType/edit',87,100,'',0,1),(97,'商品管理',3,'执行增加 商品类型','goodsType/doAdd',87,100,'',0,1),(98,'商品管理',3,'执行修改 商品类型','goodsType/doEdit',87,100,'',0,1),(99,'商品管理',3,'删除 商品类型','goodsType/delete',87,100,'',0,1),(100,'商品管理',2,'商品列表','goods',87,100,'商品列表',0,1),(101,'商品管理',3,'增加商品','goods/add',87,100,'',0,1),(102,'商品管理',3,'执行 增加商品','goods/doAdd',87,100,'',0,1),(103,'商品管理',3,'修改商品','goods/edit',87,100,'',0,1),(104,'商品管理',3,'执行 修改商品','goods/doEdit',87,100,'',0,1),(105,'商品管理',3,'删除商品','goods/delete',87,100,'',0,1),(106,'系统设置',1,'','',0,100,'',0,1),(107,'系统设置',2,'导航管理','nav',106,100,'',0,1),(108,'系统设置',3,'增加导航','nav/add',106,100,'',0,1),(109,'系统设置',3,'编辑导航','nav/edit',106,100,'',0,1),(110,'系统设置',3,'删除导航','nav/delete',106,100,'',0,1),(111,'系统设置',3,'执行增加','nav/doAdd',106,100,'',0,1),(112,'系统设置',3,'执行修改','nav/doEdit',106,100,'',0,1),(113,'系统设置',2,'商店设置','setting',106,100,'',0,1);
/*!40000 ALTER TABLE `access` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `address`
--

DROP TABLE IF EXISTS `address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `address` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `default_address` tinyint(1) DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address`
--

LOCK TABLES `address` WRITE;
/*!40000 ALTER TABLE `address` DISABLE KEYS */;
INSERT INTO `address` VALUES (42,12,'李四','15201686411','北京市 海淀区 西二旗 xxx好',0,0),(43,12,'张三','15201686411','深圳市   宝安区  xxx',0,0),(44,12,'王五','15201686411','上海市 xxx',0,0),(46,12,'王鹏朗','18243430922','吉林省',0,0),(48,12,'蔡徐坤','18243430922','1',0,0),(51,12,'我当时','18243430922','13',0,0),(52,12,'你干嘛','18243430922','艾欧',1,0);
/*!40000 ALTER TABLE `address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `focus`
--

DROP TABLE IF EXISTS `focus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `focus` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `focus_type` tinyint(1) DEFAULT NULL,
  `focus_img` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `focus`
--

LOCK TABLES `focus` WRITE;
/*!40000 ALTER TABLE `focus` DISABLE KEYS */;
INSERT INTO `focus` VALUES (14,'小米手机',1,'static/upload/20260428/1631677671.jpg','https://www.mi.com/',1115,1,1731677671),(16,'小米电视1111',1,'static/upload/20260503/1777799447.png','https://www.mi.com/',1222,0,1763167924);
/*!40000 ALTER TABLE `focus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods`
--

DROP TABLE IF EXISTS `goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cate_id` int DEFAULT NULL,
  `click_count` int DEFAULT NULL,
  `goods_number` int DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `market_price` decimal(10,2) DEFAULT NULL,
  `relation_goods` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_attr` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_gift` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_fitting` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `is_delete` tinyint DEFAULT NULL,
  `is_hot` tinyint DEFAULT NULL,
  `is_best` tinyint DEFAULT NULL,
  `is_new` tinyint DEFAULT NULL,
  `goods_type_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods`
--

LOCK TABLES `goods` WRITE;
/*!40000 ALTER TABLE `goods` DISABLE KEYS */;
INSERT INTO `goods` VALUES (19,'小米9','火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起','',38,100,1000,999.00,999.00,'20，21','','1,2,3,4,5','8GB+256GB','static/upload/20260508/1778237029987433700.png','20，21','20，21','','','<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p>圣诞节凯撒此擦hi惨白惨白此u赤壁u看</p>',0,0,0,0,4,100,1,1592392307),(20,'小米1111','火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起','',3,100,0,124124.00,124.00,'1','4','3,5','12G+512GB','static/upload/20260418/1592392495500412500.jpg','2','3','','','<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p>',0,1,1,0,3,0,1,1592392495),(21,'小米8年度旗舰222','火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起','',36,100,1000,1112.00,1113.00,'1,2','1,2','2,3,4,5','3GB+32GB','static/upload/20260501/1635849810407008900.png','1,2','1,2','1,2','1,2','<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><img src=\"http://bee.apiying.com/static/upload/20211101/1635736323217965200.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',0,1,1,1,1,11,1,1592392825),(22,'Redmi 7A','「3GB+32GB到手价仅549元」4000mAh超长续航 / 骁龙8核处理器 / 标配10W快充 / AI人脸解锁 / 大字体，大音量，无线收音机 / 整机生活防泼溅 / 极简模式，亲情守护','',2,100,1000,549.00,799.00,'','','3,4','3GB+32GB','static/upload/20260425/1592820040.jpg','','','','','<p><span style=\"color: rgb(51, 51, 51); font-family: F9ab65; font-size: 10.4922px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\">小巧机身蕴藏4000mAh大电量，配合MIUI系统级省电优化，精细调控，从此告别电量焦虑，尽情尽欢！</span></p>',0,0,1,0,1,100,1,1592820016),(23,'Redmi 智能电视 X65','全金属边框/4K超高清/MEMC运动补偿/8单元重低音音响系统','',5,100,1000,2999.00,3299.00,'','','4','56寸','static/upload/20260425/1592820111.jpg','','','','','<p><span style=\'color: rgb(176, 176, 176); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\'>全金属边框/4K超高清/MEMC运动补偿/8单元重低音音响系统</span></p>',0,0,1,0,0,100,0,1592820111),(24,'RedmiBook 13 全面屏','四窄边全面屏 / 全新十代酷睿™处理器 / 全金属超轻机身 / MX250 高性能独显 / 小米互传 / 专业「飓风」散热系统 / 11小时长续航','',20,100,1000,4499.00,4799.00,'','','4,5','8G+128G','static/upload/20260425/1592820244.jpg','','','','','<p><span style=\'color: rgb(176, 176, 176); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\'>四窄边全面屏 / 全新十代酷睿&trade;处理器 / 全金属超轻机身 / MX250 高性能独显 / 小米互传 / 专业「飓风」散热系统 / 11小时长续航</span> </p>',0,0,1,0,0,100,1,1592820244),(25,'米家电磁炉','99挡微调控火 / 支持低温烹饪 / 100+烹饪模式','',1,100,1000,299.00,399.00,'','','','','static/upload/20260425/1592820331.jpg','','','','','<p>米家电磁炉</p>',0,1,1,0,0,100,1,1592820331),(26,'黑鲨双向快充移动电源','18W双向快充 / 铠甲机身 / 一入三出 / 炫酷灯效','',37,100,1000,0.00,0.00,'','','','','static/upload/20260425/1592820494.jpg','','','','','',0,0,1,0,0,100,1,1592820494),(36,'1111','2222','',28,100,0,4444.00,5555.00,'7777','10','1,4,5','3333','static/upload/20260427/1633755416784286300.png','8888','999999999','','','',0,1,1,0,1,0,1,1783755416),(37,'1111111111','222222222222222222','',28,100,0,444444.00,5555.00,'7777777777','1000000011111100000','1,2,3,4,5','333333333333333','static/upload/20260427/1633755741820253400.png','8888888888','999999','111 1111111 111111111111111111','121212','<p>666666666666666</p>',0,1,1,1,2,0,1,1783755741),(38,'11111111214214','124214214214','',23,100,0,0.00,0.00,'','','2,5','','static/upload/20260427/1633755959396859300.png','','','','','',0,0,0,0,1,0,1,1783755959),(39,'手机','华为','',23,100,0,888.00,999.00,'','','','Mate','static/upload/20260503/1777817054796732800.png','','','','','',1,1,1,1,0,0,0,1777817054),(40,'手机','华为','',23,100,0,888.00,999.00,'','','1,2,3,4','Mate60','static/upload/20260504/1777889695886992500.jpg','','','','','<p><img src=\"/static/upload/20260504/1777889684829618500.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><img src=\"/static/upload/20260504/1777889964449012400.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><br></p>',0,1,1,1,4,0,1,1777819628),(41,'手机','华为','',27,100,0,888.00,999.00,'','','','Mate60','static/upload/20260505/1777989644429490900.png','','','','','<p><img src=\"https://java-system.oss-cn-beijing.aliyuncs.com/static/upload/20260505/1777914700338155500.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',0,1,0,0,4,0,1,1777819675),(42,'手机','华为','',23,100,0,888.00,999.00,'2','2','2,4,5','Mate60','','4','4','2','3','<p>213</p>',0,1,0,0,4,0,1,1777820813),(43,'手机','华为','',23,100,0,888.00,999.00,'','','','Mate60','','','','','','',0,1,0,0,0,0,1,1777821100),(44,'困困','华为','',24,100,0,888.00,999.00,'1','1','2,3,4','Mate60','','1','1','1','1','',0,1,0,0,5,0,1,1777887827),(45,'Apple iPhone 15','苹果官方正品手机','',28,299,88,5999.00,6299.00,'37,39','颜色:黑色,白色,粉色|容量:128G,256G','','static/upload/20250504/iphone15.png','static/upload/20260505/1777952151529022300.png','40','1,2,3','苹果,手机,iPhone','苹果iPhone15 全新正品','<p>iPhone15 搭载A16芯片，超视网膜显示屏</p>',0,1,1,1,0,10,1,1783756000),(46,'华为 Mate60 Pro','鸿蒙旗舰智能手机','HW6001',28,356,66,6499.00,6999.00,'36,40','颜色:雅丹黑,锦华白,素皮青','全网通版','static/upload/20250504/mate60.png','','38','2,3,4','华为,旗舰,鸿蒙','华为Mate60 Pro 旗舰手机','<p>卫星通话 昆仑玻璃 鸿蒙OS4.0</p>',0,1,1,1,2,9,1,1783756100),(47,'小米14 标准版','徕卡影像旗舰手机','XM1401',28,421,120,4299.00,4599.00,'41,37','颜色:黑色,白色,绿色','标准版','static/upload/20250504/xiaomi14.png','','40','1,3,5','小米,徕卡,旗舰','小米14 骁龙8Gen3 徕卡影像','<p>小米14 小屏旗舰 徕卡光学镜头</p>',0,1,0,1,2,8,1,1783756200),(48,'三星 Galaxy S24','骁龙8至尊版手机','SX2401',28,188,50,5699.00,5999.00,'','颜色:雾屿蓝,秘矿黑,月牙白','海外版','static/upload/20250504/s24.png','','','3,4,5','三星,骁龙,安卓','三星S24 超清屏幕 智能AI','<p>Galaxy S24 骁龙8至尊版 超长续航</p>',0,0,1,0,2,7,1,1783756300),(49,'vivo X100 Pro','蔡司影像旗舰','VX10001',28,266,77,4999.00,5299.00,'40,42','颜色:落日橙,白月光,辰夜黑','全网通','static/upload/20250504/x100.png','','37','1,4,5','vivo,蔡司,拍照','vivo X100 Pro 蔡司长焦镜头','<p>蔡司T*镀膜 全焦段超清影像</p>',0,1,1,0,2,6,1,1783756400),(50,'OPPO Find X7','超光影影像旗舰','OPX701',28,211,90,4699.00,4999.00,'','颜色:海阔天空,烟云紫,松影绿','标准版','static/upload/20250504/findx7.png','','40','2,4,5','OPPO,影像,旗舰','OPPO Find X7 超光影三摄','<p>超光影图像引擎 天玑9300+</p>',1,0,0,1,2,5,0,1783756500),(51,'荣耀 Magic6 Pro','青海湖电池手机','RY601',28,302,80,5199.00,5499.00,'41,43','颜色:勃朗蓝,苔原绿,祁连雪','旗舰版','static/upload/20250504/magic6.png','','41','1,2,5','荣耀,青海湖,旗舰','荣耀Magic6 Pro 超长续航','<p>青海湖双电池 骁龙8Gen3</p>',0,1,1,1,2,4,1,1783756600),(52,'红米 K70E','性价比千元旗舰','',23,568,200,1799.00,1999.00,'42,45','颜色:墨羽,晴雪,竹月','2,3','static/upload/20250504/k70e.png','','38','1,2,3','红米,性价比,千元机','红米K70E 天玑8400 大电池','<p>1.5K高光屏 5500mAh大电池</p>',0,1,0,1,0,3,1,1783756700),(53,'真我 GT Neo5 SE','144Hz电竞手机','ZM501',23,412,150,2199.00,2399.00,'','颜色:圣白幻影,最终幻想,黑色','电竞版','static/upload/20250504/gtneo5.png','','','2,3,4','真我,电竞,高刷','真我GT Neo5 SE 144Hz旗舰屏','<p>第二代骁龙7+ 100W快充</p>',0,0,0,0,1,2,1,1783756800),(54,'一加 Ace 3','超长续航游戏手机','',23,389,130,2599.00,2799.00,'44,46','颜色:鸣鹤霜,星辰黑,极光紫','2,3,4','static/upload/20250504/ace3.png','','42','1,3,4','一加,游戏,续航','一加Ace3 超大电池 游戏性能','<p>5500mAh电池 骁龙8 Gen2</p>',0,1,0,1,0,1,1,1783756900),(55,'小米手机2222','1111111111','',1,100,0,4444.00,5555.00,'7777','10','1,3,5','33332222222222','static/upload/20260502/1633944450109113300.jpg','8888','999999999','','','<p>小米10</p><p><img src=\"/static/upload/20211011/1633944295927657900.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><img src=\"/static/upload/20211011/1633944470896453500.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',0,0,0,0,2,12,1,1783755416),(56,'小米电视测试','222222222222222222','',1,100,0,444444.00,5555.00,'7777777777','1000000011111100000','1,2,3,4,5','333333333333333','static/upload/20260427/1633755741820253400.png','8888888888','999999','111 1111111 111111111111111111','121212','<p>666666666666666</p>',0,1,1,1,2,0,0,1783755741),(57,'小米手机测试111','124214214214','',1,100,0,0.00,0.00,'','','2,5','','static/upload/20260427/1633755959396859300.png','','','','','',0,1,1,1,1,0,0,1783755959),(58,'Redmi k30','6.53\"水滴大屏 | 5020mAh超长续航 | G80高性能处理器 | 全场景 AI 四摄 | 大功率扬声器 | 指纹识别 | 人脸解锁 | 红外遥控','',38,100,100,899.00,899.00,'','','','','static/upload/20260424/1637026344085801400.jpg','','','','','',0,0,0,0,0,100,1,1785502706),(59,'Xiaomi MIX 4','CUP全面屏 | 真彩原色 + 120Hz | 一体化轻量陶瓷机身 | 高通骁龙™888+ | WiFi 6 增强版 | 石墨烯「冰封」散热系统','',37,100,100,0.00,0.00,'','','','','static/upload/20260424/1637026171480899500.jpg','','','','','',0,0,0,0,0,100,1,1785503000),(60,'Xiaomi Civi','轻薄潮流设计 | 丝绒AG工艺 | 原生美肌人像 | 像素级肌肤焕新技术 | 3200万高清质感自拍 | 双柔光灯+自动对焦 | 3D曲面OLED柔性屏 | 120Hz+Dolby Vision | 4500mAh 大电量 | 55W有线闪充 | 立体声双扬声器','',36,100,100,1200.00,1400.00,'','','','','static/upload/20260424/1637026086634961500.jpg','','','','','',0,0,0,0,0,100,1,1785503077),(61,'Redmi Note 10 5G',' 5G小金刚｜旗舰长续航｜双5G待机｜5000mAh充电宝级大容量｜4800万高清相机｜天玑700八核高性能处理器','',35,100,100,0.00,0.00,'','','','','static/upload/20260424/1637025991576339600.jpg','','','','','',0,0,0,0,0,100,1,1785503644),(62,'Xiaomi 10S','骁龙870 | 对称式双扬立体声 | 1亿像素 8K电影相机 | 33W有线快充 | 30W无线快充 | 10W反向充电 | 4780mAh超大电池 | LPDDR5+UFS3.0+Wi-Fi 6 | VC液冷散热 | 双模5G','',35,100,100,2699.00,3699.00,'','','1,2,3','8GB+128GB','static/upload/20260422/1635841579767962200.jpg','','','','','<p id=\"isPasted\"><br></p><p>高通骁龙&trade;870</p><p>哈曼卡顿｜对称式双扬立体声</p><p>4780mAh 大电量</p><p>三重快充 33W有线+30W无线+10W反向充电</p><p>小至尊经典外观</p><p>LPDDR5+UFS3.0+WiFi6</p><p>1 亿像素电影相机</p><p>8K 电影模式</p><p><br></p>',0,0,1,0,1,100,1,1785841578),(63,'Xiaomi 11 Pro','至高享24期免息，赠蓝牙耳机Air2 SE，+1元得30W立式无线充','',2,100,100,0.00,0.00,'','','2,3,4','','static/upload/20260422/1635841908156579200.jpg','','','','','<p><br></p><p id=\"isPasted\" style=\"text-align: center;\"><span style=\"font-size: 24px;\">联合研发18个月</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">2亿影像投入，打造超强规格主摄</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">这是颗&ldquo;巨型大底&rdquo;的面积，甚至可以媲美专业便携式相机，超大的进光量，</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">带来了前所未有丰富的细节，&ldquo;夜视&rdquo;能力因此远超人眼，更能&ldquo;看懂&rdquo;夜色。</span></p><p><img src=\"/static/upload/20211102/1635841855622147000.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><br></p>',0,0,1,0,0,100,1,1785841907),(64,'小米移动电源3 20000mAh USB-C双向快充版','','',20,100,100,100.00,100.00,'','','','','static/upload/20260422/1635844763742258900.jpg','','','','','',0,0,0,0,0,100,1,1785844763),(65,'小米移动电源3 10000mAh 超级闪充版 （50W）','','',20,100,100,125.00,155.00,'','','','','static/upload/20260422/1635844808324401400.jpg','','','','','',0,0,0,0,0,100,1,1785844808),(66,'小米6A Type-C快充数据线','','',9,100,100,29.00,29.00,'','','','','','','','','','',0,0,0,0,0,100,1,1785845354),(67,'小米USB-C数据线 编织线版 100cm','','',9,100,100,0.00,0.00,'','','','','static/upload/20260422/1635845426055325800.jpg','','','','','<p><img src=\"/static/upload/20211102/1635845418913722200.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',0,0,0,0,0,100,1,1785845425),(68,'Redmi Note 11 Pro系列','三星AMOLED高刷屏 | JBL 对称式立体声 | 一亿像素超清影像 | 天玑920液冷芯 | VC液冷立体散热','',2,100,100,0.00,0.00,'','','','','static/upload/20260424/1637025826328576500.jpg','','','','','',0,0,0,0,0,100,1,1787025826),(69,'Valorant','华为','',3,100,100,888.00,999.00,'','','2,3,4','Mate60','static/upload/20260505/1777988848071990300.png','','','','','',0,1,0,0,4,100,1,1777988848),(70,'大电视','联想','',5,100,100,111111.00,89999.00,'','','1,2,3,4,5','','static/upload/20260505/1777990163928677800.png','','','','','',0,1,1,1,0,100,1,1777990163);
/*!40000 ALTER TABLE `goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_attr`
--

DROP TABLE IF EXISTS `goods_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods_attr` (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_id` int DEFAULT NULL,
  `attribute_cate_id` int DEFAULT NULL,
  `attribute_id` int DEFAULT NULL,
  `attribute_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `attribute_type` tinyint(1) DEFAULT NULL,
  `attribute_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods_attr`
--

LOCK TABLES `goods_attr` WRITE;
/*!40000 ALTER TABLE `goods_attr` DISABLE KEYS */;
INSERT INTO `goods_attr` VALUES (5,36,1,1,'基本信息',1,'iphone13',10,1783755417,1),(6,36,1,7,'性能	',2,'\r\n高通骁龙439八核处理器\r\n最高主频 2.0GHz\r\nAdrenoTM 505 图形处理器，最高主频 650 MHz',10,1783755417,1),(7,36,1,8,'相机',2,'1200万 AI后置相机\r\n后置1200万AI相机\r\nPDAF相位对焦\r\n人像模式，背景虚化\r\n单色温闪光灯\r\n标准 HDR\r\nAuto HDR',10,1783755417,1),(8,36,1,9,'支持蓝牙',3,'是\r\n',10,1783755417,1),(9,37,2,2,'主体',3,'111\r\n',10,1783755741,1),(10,37,2,3,'内存',1,'内存',10,1783755741,1),(11,37,2,4,'硬盘',1,'硬盘',10,1783755741,1),(12,37,2,5,'显示器',1,'显示器:',10,1783755741,1),(13,37,2,6,'支持蓝牙',3,'否',10,1783755741,1),(14,38,1,1,'基本信息',1,'124214',10,1783755959,1),(15,38,1,7,'性能	',2,'214214',10,1783755959,1),(16,38,1,8,'相机',2,'214214',10,1783755959,1),(17,38,1,9,'支持蓝牙',3,'是\r\n',10,1783755959,1),(21,42,4,10,'是否支持蓝牙',3,'是\r\n',10,1777820813,1),(22,42,4,13,'蔡徐坤',1,'2',10,1777820813,1),(23,42,4,14,'陈立农',2,'1',10,1777820813,1),(60,40,4,10,'是否支持蓝牙',3,'是\r\n',10,1777896504,1),(61,40,4,13,'蔡徐坤',1,'2',10,1777896504,1),(62,40,4,14,'陈立农',2,'1',10,1777896504,1),(63,69,4,10,'是否支持蓝牙',3,'是\r\n',10,1777988848,1),(64,69,4,13,'蔡徐坤',1,'11',10,1777988848,1),(65,69,4,14,'陈立农',2,'11',10,1777988848,1),(81,19,4,10,'是否支持蓝牙',3,'是\r\n',10,1778243435,1),(82,19,4,13,'蔡徐坤',1,'11',10,1778243435,1),(83,19,4,14,'陈立农',2,'1123',10,1778243435,1);
/*!40000 ALTER TABLE `goods_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_cate`
--

DROP TABLE IF EXISTS `goods_cate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods_cate` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `cate_img` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `template` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `pid` int DEFAULT NULL,
  `sub_title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `keywords` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `description` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `sort` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods_cate`
--

LOCK TABLES `goods_cate` WRITE;
/*!40000 ALTER TABLE `goods_cate` DISABLE KEYS */;
INSERT INTO `goods_cate` VALUES (1,'手机','','','',0,'手机','手机','手机',1,'10',1582461745),(2,'小米11 Pro pro','static/upload/20260426/1582463294.png','','',1,'小米10 Pro','小米10 Pro','小米10 Pro',0,'0',1582463294),(3,'Redmi 8','static/upload/20260426/1582463357.png','','',1,'Redmi 8 11','Redmi 8 111','Redmi 8 111',1,'11',1582463357),(4,'电视 盒子','','','',0,'电视 盒子','电视 盒子','电视 盒子',1,'0',1582463515),(5,'小米电视5 55英寸','static/upload/20260426/1582464603.png','','',4,'小米电视5 55英寸','小米电视5 55英寸','小米电视5 55英寸',1,'0',1582464603),(6,'家电 插线板','','','',0,'','','',1,'0',1582513219),(7,'出行 穿戴','','','',0,'','','',1,'0',1582513235),(8,'智能 路由器','','','',0,'','','',1,'0',1582513270),(9,'电源 配件','','','',0,'','','',1,'0',1582513285),(13,'冰箱','static/upload/20260429/1582513945.jpg','','',6,'冰箱','冰箱','冰箱',1,'0',1582513945),(14,'微波炉','static/upload/20260429/1582514001.jpg','','',6,'','','',1,'0',1582513960),(15,'小米手表','static/upload/20260429/1582514113.png','','',7,'小米手表','小米手表','小米手表',1,'0',1582514113),(16,'平衡车','static/upload/20260429/1582514151.jpg','','',7,'平衡车','平衡车','平衡车',1,'0',1582514151),(17,'路由器','static/upload/20260429/1582514289.png','','',8,'路由器','路由器','路由器',1,'0',1582514289),(18,'摄像机','static/upload/20260429/1582514318.jpg','','',8,'摄像机','摄像机','摄像机',1,'0',1582514318),(19,'全屏电视55寸','static/upload/20260429/1582514664.jpg','','',4,'','','',1,'0',1582514664),(20,'移动电源','static/upload/20260429/1582514810.png','','',9,'移动电源','移动电源','移动电源',1,'0',1582514810),(23,'手机','','','',0,'','','手机',1,'101',1731938178),(24,'电视','','','',0,'','','手机',1,'10',1731938196),(25,'笔记本 平板','','','',0,'','','手机',1,'10',1731938209),(26,'家电','','','',0,'','','手机',1,'10',1731938214),(27,'小米11Ab','static/upload/20260421/1631938291.png','','',23,'小米11立即购买','小米11','小米手机官网正品小米11推荐，小米手机小米11最新价格，有多种颜色可选，另有小米11详细介绍及图片，还有',1,'10',1731938291),(28,'Redmi 11A','static/upload/20260503/1777801325.png','http://www.itying.com','bbbb.html',23,'游戏必备','防疫','游戏必备',1,'10',1731938339),(29,'小米电视55寸JB','static/upload/20260421/1631938567.jpg','','',24,'','','',1,'10',1731938567),(30,'冰箱','static/upload/20260421/1631940993.jpg','http://www.itying.com','',26,'','','',1,'10',1731938591),(35,'Xiaomi 10S','static/upload/20260422/1635841294026066400.png','','',1,'','','',1,'10',1785817714),(36,'Xiaomi Civi','static/upload/20260422/1635841252665099500.png','','',1,'','','',1,'10',1785841252),(37,'菜鸟','static/upload/20260505/1777912568417381300.jpg','http://www.itying.com','',0,'1','1','1',1,'10',1777912568),(38,'Redmi K30S 至尊纪念版','static/upload/20260422/1635841411131518300.png','','',1,'','','',1,'10',1785841411);
/*!40000 ALTER TABLE `goods_cate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_color`
--

DROP TABLE IF EXISTS `goods_color`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods_color` (
  `id` int NOT NULL AUTO_INCREMENT,
  `color_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `color_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods_color`
--

LOCK TABLES `goods_color` WRITE;
/*!40000 ALTER TABLE `goods_color` DISABLE KEYS */;
INSERT INTO `goods_color` VALUES (1,'红色','red',1),(2,'黑色','#000',1),(3,'黄色','yellow',1),(4,'金色','#ebf10f',1),(5,'灰色','#eee',1);
/*!40000 ALTER TABLE `goods_color` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_image`
--

DROP TABLE IF EXISTS `goods_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods_image` (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_id` int DEFAULT NULL,
  `img_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `color_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods_image`
--

LOCK TABLES `goods_image` WRITE;
/*!40000 ALTER TABLE `goods_image` DISABLE KEYS */;
INSERT INTO `goods_image` VALUES (3,36,'static/upload/20260428/1633755415645620800.png',0,10,1783755417,1),(4,36,'static/upload/20260428/1633755415656163100.png',0,10,1783755417,1),(5,37,'static/upload/20260428/1633755740718752300.png',0,10,1783755741,1),(6,37,'static/upload/20260428/1633755740714630100.jpg',0,10,1783755741,1),(7,38,'static/upload/20260428/1633755956051077200.png',0,10,1783755959,1),(8,38,'static/upload/20260428/1633755956136482100.png',0,10,1783755959,1),(9,38,'static/upload/20260428/1633755956135954600.jpg',0,10,1783755959,1),(10,42,'static/upload/20260503/1777820803147764800.png',0,10,1777820813,1),(11,44,'static/upload/20260504/1777887826409470000.jpg',0,10,1777887827,1),(12,40,'static/upload/20260419/1635503037433844200.jpg',0,10,1785503038,1),(13,40,'static/upload/20260419/1635503037587034300.jpg',0,10,1785503038,1),(14,21,'static/upload/20260422/1635736448687849200.jpg',0,10,1785736455,1),(19,43,'static/upload/20260422/1635841578192734800.png',0,10,1785841580,1),(20,44,'static/upload/20260422/1635841907018281600.jpg',0,10,1785841908,1),(21,69,'static/upload/20260505/1777988827650867200.png',0,10,1777988848,1),(23,70,'static/upload/20260505/1777990163248146600.png',0,10,1777990164,1),(24,70,'static/upload/20260505/1777990163251315800.png',0,10,1777990164,1),(28,19,'static/upload/20260506/1777997192836380500.png',5,10,1777997193,1),(29,19,'static/upload/20260506/1777997332403496500.png',0,10,1777997332,1);
/*!40000 ALTER TABLE `goods_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_type`
--

DROP TABLE IF EXISTS `goods_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `status` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods_type`
--

LOCK TABLES `goods_type` WRITE;
/*!40000 ALTER TABLE `goods_type` DISABLE KEYS */;
INSERT INTO `goods_type` VALUES (4,'手机','iPhone17',1,1732299505),(5,'电视','电视',1,1732299512),(6,'笔记本','笔记本',1,1732299526),(7,'路由器','路由器',1,1732299535),(9,'衣服','衣服',1,1732361292);
/*!40000 ALTER TABLE `goods_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_type_attribute`
--

DROP TABLE IF EXISTS `goods_type_attribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goods_type_attribute` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cate_id` int DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `attr_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `attr_value` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `cate_id` (`cate_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods_type_attribute`
--

LOCK TABLES `goods_type_attribute` WRITE;
/*!40000 ALTER TABLE `goods_type_attribute` DISABLE KEYS */;
INSERT INTO `goods_type_attribute` VALUES (1,1,'基本信息','1','',1,10,1782299512),(2,2,'主体','3','111\r\n1111',1,19,1782299512),(3,2,'内存','1','',1,NULL,1782299512),(4,2,'硬盘','1','',1,NULL,1782299512),(5,2,'显示器','1','',1,111,1582361804),(6,2,'支持蓝牙','3','是\r\n否',1,1011,1582362691),(7,1,'性能	','2','',1,111,1782299512),(8,1,'相机','2','',1,0,1782299512),(9,1,'支持蓝牙','3','是\r\n否',1,0,1591844649),(10,4,'是否支持蓝牙','3','是\r\n否',1,1022,1782370943),(12,3,'尺寸1','1','',1,10,1782388221),(13,4,'蔡徐坤','1','',1,10,1777811855),(14,4,'陈立农','2','',1,10,1777811865);
/*!40000 ALTER TABLE `goods_type_attribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `manager`
--

DROP TABLE IF EXISTS `manager`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `manager` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `mobile` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `role_id` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `is_super` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `role_id` (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `manager`
--

LOCK TABLES `manager` WRITE;
/*!40000 ALTER TABLE `manager` DISABLE KEYS */;
INSERT INTO `manager` VALUES (1,'admin','e10adc3949ba59abbe56e057f20f883e','152016111','5188611114@qq.com',1,9,1777799355,1),(2,'zhangsan','e10adc3949ba59abbe56e057f20f883e','1520111122','342338691122@qq.com',1,14,1731661532,0),(6,'lisi','e10adc3949ba59abbe56e057f20f883e','1520171111','11114292@qq.com',1,16,1731156378,0);
/*!40000 ALTER TABLE `manager` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nav`
--

DROP TABLE IF EXISTS `nav`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nav` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `position` tinyint(1) DEFAULT NULL,
  `is_opennew` tinyint(1) DEFAULT NULL,
  `relation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nav`
--

LOCK TABLES `nav` WRITE;
/*!40000 ALTER TABLE `nav` DISABLE KEYS */;
INSERT INTO `nav` VALUES (1,'小米商城1','https://www.mi.com/',2,2,'21,22,23,24',10,1,1592919226),(2,'MIUI','https://www.mi.com/',1,1,'1',10,1,1592921999),(3,'小米手机','https://shouji.mi.com/',1,2,'19,20',10,1,1592922081),(4,'小米电视','https://ds.mi.com/',2,2,'39，40，41，42，43，44',10,1,1592922273),(5,'路由器','https://www.mi.com/',2,1,'25',10,1,1592922331),(8,'云服务','https://i.mi.com/',1,2,'2',10,1,1593529309),(9,'金融','https://jr.mi.com/?from=micom',1,1,'1',10,1,1593529329),(10,'有品','https://youpin.mi.com/',1,1,'1',10,1,1593529346),(11,'家电','',2,1,'1',10,1,1593529451),(12,'智能电视','',2,1,'1',10,1,1593529470),(14,'小米帮助中心','https://www.mi.com/',3,2,'12,13,14',10,1,1784788777);
/*!40000 ALTER TABLE `nav` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int DEFAULT NULL,
  `order_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `all_price` decimal(10,2) DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `zipcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `pay_status` tinyint(1) DEFAULT NULL,
  `pay_type` tinyint(1) DEFAULT NULL,
  `order_status` tinyint(1) DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
INSERT INTO `order` VALUES (39,12,'202112161333074546',4698.00,'李四','15201686411','北京市 海淀区 西二旗 xxx好11',NULL,0,0,0,1776632787),(40,12,'202112161339577439',4698.00,'王五','15201686412','上海市 xxx11 222 111',NULL,0,0,0,1776633197),(41,12,'202605092324524580',3548.00,'我当时','18243430922','13',NULL,0,0,0,1778340292),(42,12,'202605092344155851',1776.00,'王五','15201686411','上海市 xxx',NULL,0,0,0,1778341455),(43,12,'202605092358254606',888.00,'王五','15201686411','上海市 xxx',NULL,1,0,1,1778342305),(44,12,'202605101735392303',888.00,'你干嘛','18243430922','艾欧',NULL,1,0,1,1778405739),(45,12,'202605111215316146',1998.00,'你干嘛','18243430922','艾欧',NULL,0,0,0,1778472931);
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_item`
--

DROP TABLE IF EXISTS `order_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int DEFAULT NULL,
  `uid` int DEFAULT NULL,
  `product_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `product_id` int DEFAULT NULL,
  `product_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `product_price` decimal(10,2) DEFAULT NULL,
  `product_num` int DEFAULT NULL,
  `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_item`
--

LOCK TABLES `order_item` WRITE;
/*!40000 ALTER TABLE `order_item` DISABLE KEYS */;
INSERT INTO `order_item` VALUES (56,39,12,'Redmi Note 11 5G手机 1亿像素 55W有线闪充 50W无线闪充  6G+128GB 手机',20,'static/upload/20260430/1637139107685884400.jpg',3699.00,1,'6G+128GB','灰色',0),(57,39,12,'小米9-8GB+256GB',19,'static/upload/20260423/1592392307796676500.jpg',999.00,1,'8GB+256GB','红色',0),(58,40,12,'Redmi Note 11 5G手机 1亿像素 55W有线闪充 50W无线闪充  6G+128GB 手机',20,'static/upload/20260430/1637139107685884400.jpg',3699.00,1,'6G+128GB','灰色',0),(59,40,12,'小米9-8GB+256GB',19,'static/upload/20260423/1592392307796676500.jpg',999.00,1,'8GB+256GB','红色',0),(60,41,12,'小米9',19,'static/upload/20260508/1778237029987433700.png',999.00,1,'8GB+256GB','红色',0),(61,41,12,'困困',44,'',888.00,1,'Mate60','',0),(62,41,12,'小米8年度旗舰222',21,'static/upload/20260422/1635849810407008900.png',1112.00,1,'3GB+32GB','',0),(63,41,12,'Redmi 7A',22,'static/upload/20260425/1592820040.jpg',549.00,1,'3GB+32GB','',0),(64,42,12,'手机',41,'static/upload/20260505/1777989644429490900.png',888.00,2,'Mate60','',0),(65,43,12,'Valorant',69,'static/upload/20260505/1777988848071990300.png',888.00,1,'Mate60','',0),(66,44,12,'Valorant',69,'static/upload/20260505/1777988848071990300.png',888.00,1,'Mate60','',0),(67,45,12,'小米9',19,'static/upload/20260508/1778237029987433700.png',999.00,2,'8GB+256GB','红色',0);
/*!40000 ALTER TABLE `order_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (9,'超级管理员','我是一个超级管理员',1,1731072961),(14,'软件部门','软件部门',1,1731075350),(16,'销售部门','销售部门',1,1731589828);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_access`
--

DROP TABLE IF EXISTS `role_access`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_access` (
  `role_id` int NOT NULL,
  `access_id` int NOT NULL,
  KEY `role_id` (`role_id`) USING BTREE,
  KEY `access_id` (`access_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_access`
--

LOCK TABLES `role_access` WRITE;
/*!40000 ALTER TABLE `role_access` DISABLE KEYS */;
INSERT INTO `role_access` VALUES (14,52),(14,54),(14,55),(14,56),(14,57),(14,76),(14,53),(14,59),(14,60),(14,61),(14,62),(14,78),(14,79),(14,80),(14,81),(9,52),(9,54),(9,55),(9,53),(9,59),(9,60),(9,61),(9,62),(9,63),(9,64),(9,67),(9,82),(9,83),(9,84),(9,85),(9,70),(9,71),(16,53),(16,59),(16,60),(16,61),(16,62),(16,78),(16,79),(16,80),(16,81),(16,63),(16,64),(16,67),(16,82),(16,83),(16,84),(16,85);
/*!40000 ALTER TABLE `role_access` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting`
--

DROP TABLE IF EXISTS `setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `setting` (
  `id` int NOT NULL AUTO_INCREMENT,
  `site_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `site_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `site_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `site_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `no_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `site_icp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `site_tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `search_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tongji_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `appid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `app_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `end_point` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `bucket_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `oss_status` tinyint(1) DEFAULT NULL,
  `oss_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `thumbnail_size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting`
--

LOCK TABLES `setting` WRITE;
/*!40000 ALTER TABLE `setting` DISABLE KEYS */;
INSERT INTO `setting` VALUES (1,'GinShop','static/upload/20260505/1777989973990622600.png','小米','222222222','static/upload/20260505/1777912473973022300.png','2422','24','24','11111','GJoqWHXB2c9S9gwP','Lgf3weXuWITUUb17vDJfveg1jmKEe9','oss-cn-beijing.aliyuncs.com','Golang',0,'https://java-system.oss-cn-beijing.aliyuncs.com/','100,200,400');
/*!40000 ALTER TABLE `setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `last_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (12,'e10adc3949ba59abbe56e057f20f883e','12322223335','::1','',1778254874,1);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_temp`
--

DROP TABLE IF EXISTS `user_temp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_temp` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `send_count` int DEFAULT NULL,
  `add_day` int DEFAULT NULL,
  `add_time` int DEFAULT NULL,
  `sign` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_temp`
--

LOCK TABLES `user_temp` WRITE;
/*!40000 ALTER TABLE `user_temp` DISABLE KEYS */;
INSERT INTO `user_temp` VALUES (40,'::1','18243430922',1,20260508,1778250746,'f140fe7a6836c4c3d690bd58a809c56a'),(41,'::1','12322223333',3,20260508,1778252418,'2fd8fb7a18f85b796b48604529f902bd'),(42,'::1','12322223335',3,20260508,1778254827,'7f12d0e71ae8cb84476afc0b4de32571');
/*!40000 ALTER TABLE `user_temp` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-05-11 15:03:52
