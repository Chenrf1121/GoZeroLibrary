CREATE TABLE `books` (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `count` integer NOT NULL DEFAULT 0 COMMENT '书本数量',
                         `count_now` integer NOT NULL DEFAULT 0 COMMENT '当前书本数量',
                         `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '书本名称',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `number_unique` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;