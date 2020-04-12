# demo.book.com
iris mvc 完整示例站点

book_tb表语句

CREATE TABLE `book_tb` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `BookName` varchar(100) DEFAULT '' COMMENT '书名',
  `State` int(11) DEFAULT '0' COMMENT '状态',
  `Author` varchar(50) DEFAULT '' COMMENT '作者',
  `Press` varchar(100) DEFAULT '' COMMENT '出版社',
  `PublishTime` datetime DEFAULT NULL COMMENT '出版时间',
  `BookImage` varchar(300) DEFAULT '' COMMENT '图书封面',
  `Price` decimal(10,2) DEFAULT '0.00' COMMENT '售价',
  `Introduction` varchar(300) DEFAULT '' COMMENT '简介',
  `UpdateTime` datetime DEFAULT NULL,
  `AddTime` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='图书表';
