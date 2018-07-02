 CREATE TABLE `house_info` (
  `id` bigint AUTO_INCREMENT NOT NULL,
  `house_code` VARCHAR(32) CHARACTER SET ASCII COLLATE ASCII_BIN NOT NULL,
  `title` VARCHAR(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,
  `detail_url` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,
  `address` TEXT NOT NULL ,
  `total_price` VARCHAR(32) NOT NULL,
  `unit_price` VARCHAR(32) NOT NULL,
  `follow_info` VARCHAR(100)  NOT NULL ,
  `subway` VARCHAR(32) NOT NULL,
  `tax_free` VARCHAR(32) NOT NULL,
  `has_key` VARCHAR(32) NOT NULL,
   PRIMARY KEY (`id`)
 )ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;