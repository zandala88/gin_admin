/*
 Navicat Premium Data Transfer

 Source Server         : MySQLTest
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : up-admin

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 19/11/2023 10:03:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for function_basic
-- ----------------------------
DROP TABLE IF EXISTS `function_basic`;
CREATE TABLE `function_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `uri` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `sort` int NULL DEFAULT 0,
  `menu_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_function_basic_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of function_basic
-- ----------------------------
INSERT INTO `function_basic` VALUES (1, NULL, NULL, NULL, '1', 'test', '/test', 0, 1);
INSERT INTO `function_basic` VALUES (2, '2023-02-03 21:32:51.915', '2023-02-03 21:39:59.816', '2023-02-03 21:41:54.149', '1cde2a40-0b82-42e7-bc30-acf469f47c86', '修改功能测试', '/test2', 10, 1);
INSERT INTO `function_basic` VALUES (3, '2023-11-17 16:37:02.439', '2023-11-17 16:37:42.642', '2023-11-17 16:37:57.402', 'e7be48d7-d199-4d5e-8d0a-917499f82e67', '学生test', '/student/test', 0, 22);
INSERT INTO `function_basic` VALUES (5, '2023-11-17 21:45:01.720', '2023-11-17 21:45:01.720', NULL, '7e12b7df-e141-430e-8258-7eccbb025c81', '进行测试', '/ttttst', 2, 23);

-- ----------------------------
-- Table structure for menu_basic
-- ----------------------------
DROP TABLE IF EXISTS `menu_basic`;
CREATE TABLE `menu_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `parent_id` int NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `sort` int NULL DEFAULT 0,
  `path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `level` tinyint(1) NULL DEFAULT 0,
  `web_icon` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menu_basic_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu_basic
-- ----------------------------
INSERT INTO `menu_basic` VALUES (1, NULL, NULL, NULL, '1', 0, '首页', 0, '/home', 1, NULL);
INSERT INTO `menu_basic` VALUES (2, NULL, '2023-03-05 20:42:41.909', NULL, '2', 0, '设置', 0, '/setting', 0, NULL);
INSERT INTO `menu_basic` VALUES (3, NULL, NULL, NULL, '3', 2, '角色管理', 0, '/setting/role', 1, NULL);
INSERT INTO `menu_basic` VALUES (4, NULL, '2023-03-05 19:03:38.691', NULL, '4', 2, '菜单管理', 50, '/setting/menu', 1, NULL);
INSERT INTO `menu_basic` VALUES (5, '2023-02-02 21:40:52.666', '2023-02-02 21:52:07.558', '2023-02-02 21:57:21.319', '11435198-e662-40e7-987e-ab350dd4966e', 1, '修改菜单测试', 0, NULL, 0, NULL);
INSERT INTO `menu_basic` VALUES (6, '2023-02-02 21:41:28.243', '2023-02-02 21:41:28.243', '2023-02-02 21:57:21.319', 'a22472f9-c1e6-4ed8-a6f1-07f72d5cf6cc', 0, '新增菜单测试', 0, NULL, 0, NULL);
INSERT INTO `menu_basic` VALUES (7, '2023-02-02 21:42:20.668', '2023-02-02 21:42:20.668', '2023-02-02 21:57:21.319', '74189e41-4d77-425a-8017-f3deafbef233', 2, '新增菜单测试', 0, NULL, 0, NULL);
INSERT INTO `menu_basic` VALUES (8, '2023-03-05 15:55:34.591', '2023-03-05 19:03:18.842', '2023-03-05 19:07:54.819', '098c30d5-9154-41f7-b5b4-31d77f4d788f', 0, 'top3', 0, '/path', 0, NULL);
INSERT INTO `menu_basic` VALUES (9, '2023-03-05 16:17:08.379', '2023-03-05 19:02:24.270', '2023-03-05 19:07:56.981', 'b7f71057-22ec-4cb6-8a99-bc3d760dc3b5', 0, 'top2', 100, '/top', 0, NULL);
INSERT INTO `menu_basic` VALUES (10, '2023-03-05 16:23:35.611', '2023-03-05 18:49:03.459', '2023-03-05 19:07:58.963', '2aa33033-27d1-4faa-90be-22683c123ddc', 0, 'top4', 1000, '/top2', 0, NULL);
INSERT INTO `menu_basic` VALUES (11, '2023-03-05 18:47:42.482', '2023-03-05 18:50:44.520', '2023-03-05 19:07:44.578', 'f8910048-6e36-4a8b-8e02-fa40c26615d8', 0, 'top1', 1, '/top3', 0, NULL);
INSERT INTO `menu_basic` VALUES (12, '2023-03-05 19:06:45.841', '2023-03-05 19:06:45.841', '2023-03-05 19:07:52.026', '2148c1b9-f567-4414-a210-155f7ef2601f', 8, 'top3-子菜单', 0, '/top3/sub', 0, NULL);
INSERT INTO `menu_basic` VALUES (13, '2023-03-05 20:09:07.980', '2023-03-05 20:42:48.908', '2023-03-05 23:14:03.960', '8597e2b3-1994-4fc9-802f-f1535714dd4d', 0, '一号菜单', 100, '/one', 1, NULL);
INSERT INTO `menu_basic` VALUES (14, '2023-03-05 20:15:18.005', '2023-03-05 20:15:18.005', '2023-03-05 23:14:01.807', 'cc10e495-12b2-4af2-9bb4-4d01ebf9959d', 0, '二号菜单', 20, '/two', 1, NULL);
INSERT INTO `menu_basic` VALUES (15, '2023-03-05 20:17:46.095', '2023-03-05 20:41:53.471', NULL, '895132bd-106e-4705-9515-5f2bf677cb68', 13, '一号-一级', 0, '/one/one', 2, NULL);
INSERT INTO `menu_basic` VALUES (16, '2023-03-05 20:17:57.982', '2023-03-05 20:17:57.982', '2023-03-05 20:39:57.710', 'a14f6c24-e94d-4c05-8bd4-3c308024f02d', 15, '一号-二级', 0, '', 0, NULL);
INSERT INTO `menu_basic` VALUES (17, '2023-03-05 20:19:11.649', '2023-03-05 20:21:14.574', '2023-03-05 20:39:54.875', 'c8839ca6-d03d-4f27-bb61-a46e9677a765', 16, '一号-三级-按钮', 0, '', 2, NULL);
INSERT INTO `menu_basic` VALUES (18, '2023-03-05 23:42:26.523', '2023-03-05 23:46:29.207', '2023-03-05 23:55:33.417', '251603d1-6e17-46d5-bd38-ddb049286baf', 0, 'top', 0, '/top', 1, NULL);
INSERT INTO `menu_basic` VALUES (19, '2023-03-05 23:50:52.562', '2023-03-05 23:51:46.214', NULL, 'cee73782-2072-41b0-b739-d0c2afeb8d9a', 18, 'top2', 0, '/top/top2', 1, NULL);
INSERT INTO `menu_basic` VALUES (20, '2023-03-06 00:19:05.137', '2023-03-06 00:19:05.137', '2023-03-06 21:14:35.426', 'b3e52aa5-0f61-407d-9fbb-8891cdc0ff48', 0, 'top', 0, '/top', 0, NULL);
INSERT INTO `menu_basic` VALUES (21, '2023-03-06 00:19:20.323', '2023-03-06 00:19:20.323', '2023-03-06 21:14:33.645', '536d25db-30c3-4acd-b3e8-f3528a8a180f', 20, 'top2', 0, '/top/top2', 1, NULL);
INSERT INTO `menu_basic` VALUES (22, '2023-03-06 21:14:58.878', '2023-03-06 21:14:58.878', NULL, '010938c8-0b0b-4cc1-89a4-3504a3bae4cf', 2, '管理员管理', 0, '/setting/user', 1, NULL);
INSERT INTO `menu_basic` VALUES (23, '2023-11-12 11:58:58.582', '2023-11-12 11:58:58.582', NULL, '7aa8a519-0a42-4b3b-855d-8d7394761b3c', 0, '学生', 0, '/student', 0, NULL);
INSERT INTO `menu_basic` VALUES (24, '2023-11-12 11:59:15.476', '2023-11-17 16:30:51.514', '2023-11-17 16:35:52.557', '437152a8-9aa2-415a-b113-12feb0a6521a', 23, '学生新增1', 0, '/student/new', 1, NULL);
INSERT INTO `menu_basic` VALUES (25, '2023-11-17 21:49:57.776', '2023-11-17 21:49:57.776', '2023-11-17 21:50:14.918', '38e07985-e2e1-494e-8b13-54a126c4fe15', 0, '设置', 0, '/setting', 2, 'Setting');
INSERT INTO `menu_basic` VALUES (26, '2023-11-17 21:50:45.331', '2023-11-17 21:50:45.331', NULL, 'cb33cd6f-8f3a-4a19-9399-0173122418a5', 23, '学生测试', 1, '/student/test', 0, '');
INSERT INTO `menu_basic` VALUES (27, '2023-11-17 21:51:03.930', '2023-11-17 21:51:03.930', NULL, '1e9022e3-6e57-4c31-9159-aa45e2401ce3', 26, '学生修改', 1, '/student/test/botton', 2, '');

-- ----------------------------
-- Table structure for role_basic
-- ----------------------------
DROP TABLE IF EXISTS `role_basic`;
CREATE TABLE `role_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `is_admin` tinyint(1) NULL DEFAULT 0,
  `sort` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_role_basic_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_basic
-- ----------------------------
INSERT INTO `role_basic` VALUES (1, '2023-01-14 14:36:36.297', '2023-11-17 21:51:33.435', NULL, '1', '超级管理员', 1, 0);
INSERT INTO `role_basic` VALUES (2, '2023-01-14 14:36:36.297', '2023-11-17 21:54:03.924', NULL, 'd1d56591-55db-484e-96a6-d94a5a833cd9', '修改角色测试', 0, 0);
INSERT INTO `role_basic` VALUES (3, '2023-03-05 22:33:41.573', '2023-11-13 15:31:07.575', NULL, '2b4ab912-d0c0-45b1-8d5e-dbc3dfb2623f', '修改测试', 0, 0);
INSERT INTO `role_basic` VALUES (4, '2023-03-05 22:34:28.601', '2023-03-05 22:34:35.741', '2023-03-05 22:34:39.813', '7133084e-9a22-4291-9c14-379281e01ebf', '超管2', 0, 0);
INSERT INTO `role_basic` VALUES (6, '2023-11-13 15:33:08.220', '2023-11-13 15:34:10.865', NULL, '139dc6c7-f215-4186-a698-0c856ca612d5', '修改测试', 0, 0);

-- ----------------------------
-- Table structure for role_function
-- ----------------------------
DROP TABLE IF EXISTS `role_function`;
CREATE TABLE `role_function`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_id` int NULL DEFAULT NULL,
  `function_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_role_function_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_function
-- ----------------------------
INSERT INTO `role_function` VALUES (1, NULL, NULL, NULL, 1, 1);
INSERT INTO `role_function` VALUES (2, '2023-01-14 14:36:36.320', '2023-01-14 14:36:36.320', '2023-01-27 21:19:57.182', 2, 1);
INSERT INTO `role_function` VALUES (3, '2023-01-27 21:19:57.187', '2023-01-27 21:19:57.187', '2023-01-28 22:17:29.078', 2, 1);
INSERT INTO `role_function` VALUES (4, '2023-01-28 22:17:29.080', '2023-01-28 22:17:29.080', '2023-01-28 22:17:47.611', 2, 1);
INSERT INTO `role_function` VALUES (5, '2023-01-28 22:17:47.613', '2023-01-28 22:17:47.613', '2023-01-28 22:18:13.898', 2, 1);
INSERT INTO `role_function` VALUES (6, '2023-01-28 22:18:13.901', '2023-01-28 22:18:13.901', NULL, 2, 1);

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_id` int NULL DEFAULT NULL,
  `menu_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_role_menu_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 73 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (1, NULL, NULL, '2023-11-17 21:51:33.435', 1, 1);
INSERT INTO `role_menu` VALUES (2, NULL, NULL, '2023-11-17 21:51:33.435', 1, 2);
INSERT INTO `role_menu` VALUES (3, NULL, NULL, '2023-11-17 21:51:33.435', 1, 3);
INSERT INTO `role_menu` VALUES (27, '2023-03-06 00:17:26.951', '2023-03-06 00:17:26.951', '2023-03-06 00:18:29.182', 2, 2);
INSERT INTO `role_menu` VALUES (28, '2023-03-06 00:18:29.184', '2023-03-06 00:18:29.184', '2023-11-12 11:58:07.522', 2, 1);
INSERT INTO `role_menu` VALUES (29, '2023-03-06 00:18:29.184', '2023-03-06 00:18:29.184', '2023-11-12 11:58:07.522', 2, 2);
INSERT INTO `role_menu` VALUES (30, '2023-03-06 00:18:29.184', '2023-03-06 00:18:29.184', '2023-11-12 11:58:07.522', 2, 3);
INSERT INTO `role_menu` VALUES (31, '2023-11-12 11:58:07.523', '2023-11-12 11:58:07.523', '2023-11-12 11:59:28.401', 2, 1);
INSERT INTO `role_menu` VALUES (32, '2023-11-12 11:58:07.523', '2023-11-12 11:58:07.523', '2023-11-12 11:59:28.401', 2, 2);
INSERT INTO `role_menu` VALUES (33, '2023-11-12 11:58:07.523', '2023-11-12 11:58:07.523', '2023-11-12 11:59:28.401', 2, 3);
INSERT INTO `role_menu` VALUES (34, '2023-11-12 11:58:07.523', '2023-11-12 11:58:07.523', '2023-11-12 11:59:28.401', 2, 4);
INSERT INTO `role_menu` VALUES (35, '2023-11-12 11:59:28.402', '2023-11-12 11:59:28.402', '2023-11-12 11:59:31.023', 2, 1);
INSERT INTO `role_menu` VALUES (36, '2023-11-12 11:59:28.402', '2023-11-12 11:59:28.402', '2023-11-12 11:59:31.023', 2, 3);
INSERT INTO `role_menu` VALUES (37, '2023-11-12 11:59:28.402', '2023-11-12 11:59:28.402', '2023-11-12 11:59:31.023', 2, 4);
INSERT INTO `role_menu` VALUES (38, '2023-11-12 11:59:28.402', '2023-11-12 11:59:28.402', '2023-11-12 11:59:31.023', 2, 23);
INSERT INTO `role_menu` VALUES (39, '2023-11-12 11:59:31.023', '2023-11-12 11:59:31.023', '2023-11-17 21:52:01.743', 2, 1);
INSERT INTO `role_menu` VALUES (40, '2023-11-12 11:59:31.023', '2023-11-12 11:59:31.023', '2023-11-17 21:52:01.743', 2, 3);
INSERT INTO `role_menu` VALUES (41, '2023-11-12 11:59:31.023', '2023-11-12 11:59:31.023', '2023-11-17 21:52:01.743', 2, 4);
INSERT INTO `role_menu` VALUES (42, '2023-11-12 11:59:31.023', '2023-11-12 11:59:31.023', '2023-11-17 21:52:01.743', 2, 23);
INSERT INTO `role_menu` VALUES (43, '2023-11-12 11:59:31.023', '2023-11-12 11:59:31.023', '2023-11-17 21:52:01.743', 2, 24);
INSERT INTO `role_menu` VALUES (44, '2023-11-13 15:31:07.576', '2023-11-13 15:31:07.576', NULL, 3, 19);
INSERT INTO `role_menu` VALUES (45, '2023-11-13 15:33:08.221', '2023-11-13 15:33:08.221', '2023-11-13 15:33:50.115', 6, 19);
INSERT INTO `role_menu` VALUES (46, '2023-11-13 15:33:30.578', '2023-11-13 15:33:30.578', NULL, 0, 19);
INSERT INTO `role_menu` VALUES (47, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 1);
INSERT INTO `role_menu` VALUES (48, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 2);
INSERT INTO `role_menu` VALUES (49, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 3);
INSERT INTO `role_menu` VALUES (50, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 4);
INSERT INTO `role_menu` VALUES (51, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 22);
INSERT INTO `role_menu` VALUES (52, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 23);
INSERT INTO `role_menu` VALUES (53, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 26);
INSERT INTO `role_menu` VALUES (54, '2023-11-17 21:51:33.437', '2023-11-17 21:51:33.437', NULL, 1, 27);
INSERT INTO `role_menu` VALUES (55, '2023-11-17 21:52:01.743', '2023-11-17 21:52:01.743', '2023-11-17 21:53:50.799', 2, 1);
INSERT INTO `role_menu` VALUES (56, '2023-11-17 21:52:01.743', '2023-11-17 21:52:01.743', '2023-11-17 21:53:50.799', 2, 3);
INSERT INTO `role_menu` VALUES (57, '2023-11-17 21:52:01.743', '2023-11-17 21:52:01.743', '2023-11-17 21:53:50.799', 2, 4);
INSERT INTO `role_menu` VALUES (58, '2023-11-17 21:52:01.743', '2023-11-17 21:52:01.743', '2023-11-17 21:53:50.799', 2, 23);
INSERT INTO `role_menu` VALUES (59, '2023-11-17 21:52:01.743', '2023-11-17 21:52:01.743', '2023-11-17 21:53:50.799', 2, 26);
INSERT INTO `role_menu` VALUES (60, '2023-11-17 21:52:01.743', '2023-11-17 21:52:01.743', '2023-11-17 21:53:50.799', 2, 27);
INSERT INTO `role_menu` VALUES (61, '2023-11-17 21:53:50.800', '2023-11-17 21:53:50.800', '2023-11-17 21:54:03.925', 2, 1);
INSERT INTO `role_menu` VALUES (62, '2023-11-17 21:53:50.800', '2023-11-17 21:53:50.800', '2023-11-17 21:54:03.925', 2, 3);
INSERT INTO `role_menu` VALUES (63, '2023-11-17 21:53:50.800', '2023-11-17 21:53:50.800', '2023-11-17 21:54:03.925', 2, 4);
INSERT INTO `role_menu` VALUES (64, '2023-11-17 21:53:50.800', '2023-11-17 21:53:50.800', '2023-11-17 21:54:03.925', 2, 23);
INSERT INTO `role_menu` VALUES (65, '2023-11-17 21:53:50.800', '2023-11-17 21:53:50.800', '2023-11-17 21:54:03.925', 2, 26);
INSERT INTO `role_menu` VALUES (66, '2023-11-17 21:53:50.800', '2023-11-17 21:53:50.800', '2023-11-17 21:54:03.925', 2, 27);
INSERT INTO `role_menu` VALUES (67, '2023-11-17 21:54:03.925', '2023-11-17 21:54:03.925', NULL, 2, 1);
INSERT INTO `role_menu` VALUES (68, '2023-11-17 21:54:03.925', '2023-11-17 21:54:03.925', NULL, 2, 3);
INSERT INTO `role_menu` VALUES (69, '2023-11-17 21:54:03.925', '2023-11-17 21:54:03.925', NULL, 2, 4);
INSERT INTO `role_menu` VALUES (70, '2023-11-17 21:54:03.925', '2023-11-17 21:54:03.925', NULL, 2, 23);
INSERT INTO `role_menu` VALUES (71, '2023-11-17 21:54:03.925', '2023-11-17 21:54:03.925', NULL, 2, 26);
INSERT INTO `role_menu` VALUES (72, '2023-11-17 21:54:03.925', '2023-11-17 21:54:03.925', NULL, 2, 27);

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `username` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `password` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `role_identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_basic_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
INSERT INTO `user_basic` VALUES (1, '2023-02-21 22:17:59.365', '2023-02-21 22:17:59.365', NULL, '1', 'get', 'e10adc3949ba59abbe56e057f20f883e', '13333333333', '1', NULL);
INSERT INTO `user_basic` VALUES (2, '2023-02-21 22:17:59.365', '2023-11-12 14:18:23.578', NULL, '2', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '13344445555', 'd1d56591-55db-484e-96a6-d94a5a833cd9', NULL);
INSERT INTO `user_basic` VALUES (5, '2023-03-06 23:05:51.427', '2023-03-06 23:05:51.427', NULL, '4285b6fa-0559-4274-8b19-4c78e4d04df7', 'get1', 'e10adc3949ba59abbe56e057f20f883e', '', '1', NULL);
INSERT INTO `user_basic` VALUES (6, '2023-11-14 12:31:50.214', '2023-11-14 12:35:57.336', '2023-11-14 12:38:50.008', 'a3053ad4-9376-4e29-b0f3-bdf8463c76dc', 'get23', 'e10adc3949ba59abbe56e057f20f883e', '110', '2b4ab912-d0c0-45b1-8d5e-dbc3dfb2623f', NULL);

SET FOREIGN_KEY_CHECKS = 1;
