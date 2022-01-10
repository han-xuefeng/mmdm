/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : mmdm

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 11/01/2022 00:25:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mmdm_admin
-- ----------------------------
DROP TABLE IF EXISTS `mmdm_admin`;
CREATE TABLE `mmdm_admin`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '盐',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mmdm_admin
-- ----------------------------
INSERT INTO `mmdm_admin` VALUES (1, 'admin', 'admin', '2823d896e9822c0833d41d4904f0c00756d718570fce49b9a379a62c804689d3', '2020-04-10 16:42:05', '2022-01-11 00:12:35', 0);

-- ----------------------------
-- Table structure for mmdm_datagroup
-- ----------------------------
DROP TABLE IF EXISTS `mmdm_datagroup`;
CREATE TABLE `mmdm_datagroup`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `source_id` bigint NOT NULL COMMENT '数据源id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '组名',
  `prefix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '组前缀规则',
  `group_config` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '组命令执行前的sql',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '数据源表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mmdm_datagroup
-- ----------------------------
INSERT INTO `mmdm_datagroup` VALUES (11, 7, '测试项目1', 'test_', '', '2021-12-28 23:11:23', '2021-12-28 23:11:23', 0);

-- ----------------------------
-- Table structure for mmdm_datasource
-- ----------------------------
DROP TABLE IF EXISTS `mmdm_datasource`;
CREATE TABLE `mmdm_datasource`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `host` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'host',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `port` bigint NOT NULL COMMENT '端口',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '数据源表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mmdm_datasource
-- ----------------------------
INSERT INTO `mmdm_datasource` VALUES (7, '这是测试环境1', '192.168.1.1', 'root', 3306, 'root', '2021-12-27 23:25:16', '2021-12-27 23:25:16', 0);
INSERT INTO `mmdm_datasource` VALUES (8, '这是测试环境2', '192.168.1.1', 'root', 3306, 'root', '2021-12-27 23:25:20', '2021-12-27 23:25:20', 0);
INSERT INTO `mmdm_datasource` VALUES (9, '这是测试环境3', '192.168.1.1', 'root', 3306, 'root', '2021-12-27 23:25:23', '2021-12-27 23:25:23', 0);

-- ----------------------------
-- Table structure for mmdm_sql_task
-- ----------------------------
DROP TABLE IF EXISTS `mmdm_sql_task`;
CREATE TABLE `mmdm_sql_task`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务名',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务描述',
  `group_id` bigint NOT NULL COMMENT '组id',
  `sql` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '需要执行的sql',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0默认 1执行中 2执行结束',
  `exec_begin` bigint NULL DEFAULT 0 COMMENT '命令执行开始时间',
  `exec_over` bigint NULL DEFAULT 0 COMMENT '命令执行结束时间',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '数据源表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mmdm_sql_task
-- ----------------------------
INSERT INTO `mmdm_sql_task` VALUES (16, '我要执行sql123', 'test_', 11, '111', 1, 0, 0, '2021-12-29 23:18:21', '2021-12-29 23:26:12', 0);

SET FOREIGN_KEY_CHECKS = 1;
