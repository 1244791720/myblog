CREATE TABLE blog.article(
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `type_id` int NOT NULL,
                             `article_title` varchar(255) NOT NULL DEFAULT '',
                             `article_content` varchar(255) NOT NULL DEFAULT '',
                             `article_author` varchar(255) NOT NULL DEFAULT '原创',
                             `view_num` int NOT NULL DEFAULT '0',
                             `comment_num` int NOT NULL DEFAULT '0',
                             `like_num` int NOT NULL DEFAULT '0',
                             `cover_url` varchar(255) NOT NULL DEFAULT '',
                             `is_del` int(1) NOT NULL DEFAULT '0',
                             `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;