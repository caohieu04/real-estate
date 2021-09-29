DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`(
  `user_id` INT NOT NULL,
  `decoration_id`INT NOT NULL,
  `quantity` INT NOT NULL,
  `status` INT NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `decoration_id`),
  KEY `decoration_id` (`decoration_id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `icon` json DEFAULT NULL,
  `status` INT NOT NULL default '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `status` INT NOT NULL default '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `decoration_likes`;
CREATE TABLE `decoration_likes` (
  `user_id` INT NOT NULL,
  `decoration_id` INT NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `decoration_id`)
  KEY `decoration_id` (`decoration_id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `decoration_ratings`;
CREATE TABLE `decoration_ratings` (
  `id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `decoration_id` INT NOT NULL,
  `poINT` float default '0',
  `comment` text,
  `status` INT NOT NULL default '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `decoration_id` (`decoration_id`) USING BTREE
) ENGINE=InnoDB;
