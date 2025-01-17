/*
 Navicat Premium Data Transfer

 Source Server         : SERVER_LARAGON
 Source Server Type    : MySQL
 Source Server Version : 50724 (5.7.24)
 Source Host           : 127.0.0.1:3306
 Source Schema         : golang_test_sigma

 Target Server Type    : MySQL
 Target Server Version : 50724 (5.7.24)
 File Encoding         : 65001

 Date: 18/01/2025 00:31:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_perusahaan
-- ----------------------------
DROP TABLE IF EXISTS `tb_perusahaan`;
CREATE TABLE `tb_perusahaan`  (
  `perusahaan_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `perusahaan_nama` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `perusahaan_fee` int(11) NULL DEFAULT NULL,
  `perusahaan_alamat` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `perusahaan_create_at` datetime NULL DEFAULT NULL,
  `perusahaan_update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`perusahaan_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_perusahaan
-- ----------------------------
INSERT INTO `tb_perusahaan` VALUES ('6e323906d6f44721eb79183cbddb33b9', 'PT ABC', 25000, 'Bali', '2025-01-17 08:11:39', '2025-01-17 08:11:39');
INSERT INTO `tb_perusahaan` VALUES ('9168f757a958b2fe74bfc3f476e7afff', 'PT XYZ', 5000, 'Jakarta', '2025-01-17 08:10:19', '2025-01-17 08:10:19');

-- ----------------------------
-- Table structure for tb_perusahaan_asset
-- ----------------------------
DROP TABLE IF EXISTS `tb_perusahaan_asset`;
CREATE TABLE `tb_perusahaan_asset`  (
  `perusahaan_asset_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `perusahaan_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `perusahaan_asset_nama` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `perusahaan_asset_otr_price` int(11) NULL DEFAULT NULL,
  `perusahaan_asset_stock_availability` mediumint(6) NULL DEFAULT NULL,
  `perusahaan_asset_create_at` datetime NULL DEFAULT NULL,
  `perusahaan_asset_update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`perusahaan_asset_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_perusahaan_asset
-- ----------------------------
INSERT INTO `tb_perusahaan_asset` VALUES ('2e9c3e91ae9703703ab4956db442c7cf', '9168f757a958b2fe74bfc3f476e7afff', 'Mobil', 1000000, 1, '2025-01-17 11:36:39', '2025-01-17 11:36:39');

-- ----------------------------
-- Table structure for tb_tenor
-- ----------------------------
DROP TABLE IF EXISTS `tb_tenor`;
CREATE TABLE `tb_tenor`  (
  `tenor_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `tenor` int(11) NULL DEFAULT NULL,
  `tenor_max_limit` int(11) NULL DEFAULT NULL,
  `tenor_interest` int(11) NULL DEFAULT NULL,
  `tenor_create_at` datetime NULL DEFAULT NULL,
  `tenor_update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`tenor_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_tenor
-- ----------------------------
INSERT INTO `tb_tenor` VALUES ('065bf2a30e881b1757f4246dc9734570', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9', 6, 2000000, 20, '2025-01-16 16:40:08', '2025-01-16 16:40:08');
INSERT INTO `tb_tenor` VALUES ('5eb34323b4d5c7bdb24f9002d495844c', 'e9f77b213b7708fc51b095eedaae467c', 3, 500000, 15, '2025-01-16 16:29:43', '2025-01-16 16:29:43');
INSERT INTO `tb_tenor` VALUES ('6f7fd97eec62c2046233556ffab421e9', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9', 2, 1200000, 10, '2025-01-16 16:39:38', '2025-01-16 16:39:38');
INSERT INTO `tb_tenor` VALUES ('aa9702e9aa02823ef96a4a50c53b7958', 'e9f77b213b7708fc51b095eedaae467c', 6, 700000, 20, '2025-01-16 16:30:13', '2025-01-16 16:30:13');
INSERT INTO `tb_tenor` VALUES ('b5feb429ad0fdb74c946472e99835703', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9', 3, 1500000, 15, '2025-01-16 16:39:50', '2025-01-16 16:39:50');
INSERT INTO `tb_tenor` VALUES ('ddaaeb7b6bae734173bb942b69836e7f', 'e9f77b213b7708fc51b095eedaae467c', 1, 100000, 5, '2025-01-16 14:15:52', '2025-01-16 14:15:52');
INSERT INTO `tb_tenor` VALUES ('fc4f7247e20ee1d179b219286b558329', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9', 1, 1000000, 5, '2025-01-16 16:39:18', '2025-01-16 16:39:18');

-- ----------------------------
-- Table structure for tb_transaction
-- ----------------------------
DROP TABLE IF EXISTS `tb_transaction`;
CREATE TABLE `tb_transaction`  (
  `transaction_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `transaction_user_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `transaction_tenor_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `transaction_perusahaan_asset_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `transaction_create_at` datetime NULL DEFAULT NULL,
  `transaction_update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`transaction_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_transaction
-- ----------------------------
INSERT INTO `tb_transaction` VALUES ('dd743770f6c1373be634ae63d87e7e3e', 'affe0a6a8e58f36bab0dcc1cb8bbc8c9', 'fc4f7247e20ee1d179b219286b558329', '2e9c3e91ae9703703ab4956db442c7cf', '2025-01-17 17:42:21', '2025-01-17 17:42:21');

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `user_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `nik` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `full_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `legal_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `tempat_lahir` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `tanggal_lahir` date NULL DEFAULT NULL,
  `gaji` int(30) NULL DEFAULT NULL,
  `foto_ktp` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `foto_ktp_path` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `foto_selfie` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `foto_selfie_path` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `user_create_at` datetime NULL DEFAULT NULL,
  `user_update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES ('affe0a6a8e58f36bab0dcc1cb8bbc8c9', '143234343', 'Annisa Anggraini', 'Annisa', 'kediri', '1998-08-02', 10000000, '96d2a5e4ce8063f5254836a3503b1885.jpg', 'uploaded_files/', '871043649298f9e9151dd68f6474330e.png', '', '2025-01-16 16:32:53', '2025-01-16 16:32:53');
INSERT INTO `tb_user` VALUES ('e9f77b213b7708fc51b095eedaae467c', '243234344', 'Budi Santoso', 'Budi', 'kediri', '1998-08-02', 500000, 'e1de528a48bcb4dd58e324bf39837667.jpg', 'uploaded_files/', '12bb1c7a74c814a7633cbbecc8abb669.png', '', '2025-01-16 13:01:43', '2025-01-16 13:01:43');

SET FOREIGN_KEY_CHECKS = 1;
