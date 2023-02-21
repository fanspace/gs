测试用grpc-server, 使用go-spring, grp starter, xorm,
<pre><code>
DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
`id` bigint NOT NULL AUTO_INCREMENT,
`title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
`content` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci,
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4566 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
</pre></code>
