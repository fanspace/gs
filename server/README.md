测试用grpc-server, 使用go-spring, grpc starter,xorm
<pre><code>
DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`showname` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
`email` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `id_UNIQUE` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=65737 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
</pre></code>
