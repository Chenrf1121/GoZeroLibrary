CREATE TABLE `borrow` (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '账号',
                         `book_id` integer NOT NULL DEFAULT 0 COMMENT '书本号',
                         `isreturn` integer not null default 0 comment '是否归还，0未归还，1归还。',
                         `borrow_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `return_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;