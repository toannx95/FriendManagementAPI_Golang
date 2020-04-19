CREATE DATABASE IF NOT EXISTS `testdb`;

USE `testdb`;

START TRANSACTION;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(225) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `relationship`;
CREATE TABLE `relationship` (
  `id` int NOT NULL AUTO_INCREMENT,
  `first_email_id` int DEFAULT NULL,
  `second_email_id` int DEFAULT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;