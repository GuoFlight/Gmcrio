DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `username` varchar(40) NOT NULL UNIQUE,
    `pwd` varchar(70) NOT NULL,
    `disable` BOOLEAN DEFAULT FALSE,
    `comment` TEXT DEFAULT ''
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- 此密码是123456
-- insert into user(username,pwd) values('user1','$2a$12$qlq7rDGdaFr/LY1BKO7SPeHrMPled26kL9HTztBdKtii4waYBBRGi');