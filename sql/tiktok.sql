/*
 Navicat MySQL Data Transfer

 Source Server         : a
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 08/08/2023 23:05:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL,
  `video_id` bigint NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_videos_comments`(`video_id` ASC) USING BTREE,
  INDEX `fk_users_comments`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_videos_comments` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 542 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (511, 1, 1, 'aa', '2023-02-08 22:59:25.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (512, 2, 2, 'bb', '2023-08-01 23:01:28.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (513, 1, 3, 'cc', '2023-08-01 23:01:32.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (514, 3, 4, 'dd', '2023-08-08 23:01:35.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (515, 2, 5, 'ee', '2023-08-08 23:01:39.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (516, 4, 6, 'fgf', '2023-08-08 23:01:41.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (517, 3, 7, 'aa', '2023-08-08 23:01:43.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (518, 5, 8, 'dad', '2023-08-08 23:01:45.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (519, 6, 9, '太好了', '2023-08-08 23:01:48.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (520, 3, 10, '太好了', '2023-08-08 23:01:50.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (521, 2, 11, '太好了', '2023-08-08 23:01:52.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (522, 7, 12, '太好了', '2023-08-08 23:01:54.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (523, 6, 13, '太好了', '2023-08-08 23:01:58.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (524, 8, 14, '太好了', '2023-08-08 23:02:00.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (525, 5, 15, '太好了', '2023-08-08 23:02:05.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (526, 9, 16, '太好了', '2023-08-08 23:02:06.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (527, 8, 17, '太好了', '2023-08-08 23:02:08.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (528, 10, 18, '太好了', '2023-08-08 23:02:10.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (529, 9, 19, '太好了', '2023-08-08 23:02:12.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (530, 10, 20, '太好了', '2023-08-08 23:02:14.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (531, 1, 21, '太好了', '2023-08-08 23:02:16.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (532, 10, 22, '太好了', '2023-08-08 23:02:18.000', '2022-02-28 10:23:52.000');
INSERT INTO `comments` VALUES (533, 10, 23, '太好了', '2023-08-08 23:02:21.000', '2022-02-28 10:23:52.000');

-- ----------------------------
-- Table structure for user_favor_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_favor_videos`;
CREATE TABLE `user_favor_videos`  (
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  PRIMARY KEY (`user_id`, `video_id`) USING BTREE,
  INDEX `fk_user_favor_videos_video`(`video_id` ASC) USING BTREE,
  CONSTRAINT `fk_user_favor_videos_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_favor_videos_video` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_favor_videos
-- ----------------------------
INSERT INTO `user_favor_videos` VALUES (1, 1);
INSERT INTO `user_favor_videos` VALUES (2, 2);
INSERT INTO `user_favor_videos` VALUES (1, 3);
INSERT INTO `user_favor_videos` VALUES (3, 4);
INSERT INTO `user_favor_videos` VALUES (2, 5);
INSERT INTO `user_favor_videos` VALUES (4, 6);
INSERT INTO `user_favor_videos` VALUES (3, 7);
INSERT INTO `user_favor_videos` VALUES (5, 8);
INSERT INTO `user_favor_videos` VALUES (6, 9);
INSERT INTO `user_favor_videos` VALUES (3, 10);
INSERT INTO `user_favor_videos` VALUES (2, 11);
INSERT INTO `user_favor_videos` VALUES (7, 12);
INSERT INTO `user_favor_videos` VALUES (6, 13);
INSERT INTO `user_favor_videos` VALUES (8, 14);
INSERT INTO `user_favor_videos` VALUES (5, 15);
INSERT INTO `user_favor_videos` VALUES (9, 16);
INSERT INTO `user_favor_videos` VALUES (8, 17);
INSERT INTO `user_favor_videos` VALUES (10, 18);
INSERT INTO `user_favor_videos` VALUES (9, 19);
INSERT INTO `user_favor_videos` VALUES (10, 20);
INSERT INTO `user_favor_videos` VALUES (1, 21);
INSERT INTO `user_favor_videos` VALUES (10, 22);
INSERT INTO `user_favor_videos` VALUES (10, 23);

-- ----------------------------
-- Table structure for user_logins
-- ----------------------------
DROP TABLE IF EXISTS `user_logins`;
CREATE TABLE `user_logins`  (
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  CONSTRAINT `fk_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_logins
-- ----------------------------
INSERT INTO `user_logins` VALUES ('john.doe', '123456', 1);
INSERT INTO `user_logins` VALUES ('jane.smith', 'password123', 2);
INSERT INTO `user_logins` VALUES ('michael.johnson', 'qwerty', 3);
INSERT INTO `user_logins` VALUES ('emily.davis', 'pass123', 4);
INSERT INTO `user_logins` VALUES ('david.brown', 'abc123', 5);
INSERT INTO `user_logins` VALUES ('sarah.wilson', 'password', 6);
INSERT INTO `user_logins` VALUES ('robert.taylor', 'securepass', 7);
INSERT INTO `user_logins` VALUES ('jennifer.anderson', 'passw0rd', 8);
INSERT INTO `user_logins` VALUES ('william.martinez', '321pass', 9);
INSERT INTO `user_logins` VALUES ('jessica.moore', 'mypassword', 10);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `follow_count` bigint NULL DEFAULT NULL,
  `follower_count` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1021 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'John Doe', 100, 200);
INSERT INTO `users` VALUES (2, 'Jane Smith', 50, 150);
INSERT INTO `users` VALUES (3, 'Michael Johnson', 300, 500);
INSERT INTO `users` VALUES (4, 'Emily Davis', 70, 120);
INSERT INTO `users` VALUES (5, 'David Brown', 150, 250);
INSERT INTO `users` VALUES (6, 'Sarah Wilson', 80, 180);
INSERT INTO `users` VALUES (7, 'Robert Taylor', 200, 400);
INSERT INTO `users` VALUES (8, 'Jennifer Anderson', 90, 220);
INSERT INTO `users` VALUES (9, 'William Martinez', 180, 300);
INSERT INTO `users` VALUES (10, 'Jessica Moore', 120, 280);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL DEFAULT NULL,
  `play_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `cover_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `favorite_count` bigint NULL DEFAULT NULL,
  `comment_count` bigint NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_users_videos`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_users_videos` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, 1, 'http://playurl1', 'http://cover1', 50, 20, 'Video 1', 1656789432);
INSERT INTO `videos` VALUES (2, 2, 'http://playurl2', 'http://cover2', 30, 10, 'Video 2', 1656789435);
INSERT INTO `videos` VALUES (3, 1, 'http://playurl3', 'http://cover3', 80, 15, 'Video 3', 1656789440);
INSERT INTO `videos` VALUES (4, 3, 'http://playurl4', 'http://cover4', 60, 25, 'Video 4', 1656789450);
INSERT INTO `videos` VALUES (5, 2, 'http://playurl5', 'http://cover5', 40, 30, 'Video 5', 1656789455);
INSERT INTO `videos` VALUES (6, 4, 'http://playurl6', 'http://cover6', 70, 18, 'Video 6', 1656789460);
INSERT INTO `videos` VALUES (7, 3, 'http://playurl7', 'http://cover7', 90, 12, 'Video 7', 1656789470);
INSERT INTO `videos` VALUES (8, 5, 'http://playurl8', 'http://cover8', 55, 22, 'Video 8', 1656789475);
INSERT INTO `videos` VALUES (9, 6, 'http://playurl9', 'http://cover9', 35, 28, 'Video 9', 1656789480);
INSERT INTO `videos` VALUES (10, 3, 'http://playurl10', 'http://cover10', 45, 16, 'Video 10', 1656789490);
INSERT INTO `videos` VALUES (11, 2, 'http://playurl11', 'http://cover11', 75, 21, 'Video 11', 1656789495);
INSERT INTO `videos` VALUES (12, 7, 'http://playurl12', 'http://cover12', 65, 33, 'Video 12', 1656789500);
INSERT INTO `videos` VALUES (13, 6, 'http://playurl13', 'http://cover13', 85, 14, 'Video 13', 1656789510);
INSERT INTO `videos` VALUES (14, 8, 'http://playurl14', 'http://cover14', 53, 27, 'Video 14', 1656789515);
INSERT INTO `videos` VALUES (15, 5, 'http://playurl15', 'http://cover15', 43, 19, 'Video 15', 1656789520);
INSERT INTO `videos` VALUES (16, 9, 'http://playurl16', 'http://cover16', 63, 31, 'Video 16', 1656789530);
INSERT INTO `videos` VALUES (17, 8, 'http://playurl17', 'http://cover17', 93, 17, 'Video 17', 1656789535);
INSERT INTO `videos` VALUES (18, 10, 'http://playurl18', 'http://cover18', 57, 23, 'Video 18', 1656789540);
INSERT INTO `videos` VALUES (19, 9, 'http://playurl19', 'http://cover19', 37, 29, 'Video 19', 1656789550);
INSERT INTO `videos` VALUES (20, 10, 'http://playurl20', 'http://cover20', 47, 15, 'Video 20', 1656789560);
INSERT INTO `videos` VALUES (21, 1, 'http://playurl21', 'http://cover21', 77, 24, 'Video 21', 1656789565);
INSERT INTO `videos` VALUES (22, 10, 'http://playurl22', 'http://cover22', 67, 32, 'Video 22', 1656789570);
INSERT INTO `videos` VALUES (23, 10, 'http://playurl23', 'http://cover23', 87, 18, 'Video 23', 1656789580);

SET FOREIGN_KEY_CHECKS = 1;
