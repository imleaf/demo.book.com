# demo.book.com
iris mvc 完整示例站点

项目结构

```
demo.book.com
—— conf                           //配置相关的放在此目录
    —— sysconf.go                 //配置读取代码
    —— web.config                 //自定义文本配置
—— dao
    —— book_dao.go                //book表的xorm操作代码
—— dbsource
    —— dbsource.go                //数据引擎，单例模式
—— log
    —— 2020
        —— 04
            —— 2020-04-12.log     //站点日志，按日期存储
—— models
    —— goxorm                     //非站点使用，用于生成映射类的工具
    —— book_tb.go                 //xorm根据数据库自动生成的映射类
—— services
    —— book_service.go            //业务操作类，调用的是book_dao.go，可以根据实际业务增加额外代码，例如从缓存读取数据等。
—— web
    —— content                    //存放站点静态资源，css、image、js等
        —— mould
            —— book  
    —— controllers
        —— BookController.go        //书城业务控制器
        —— DemoController.go        //一些辅助demo控制器，测试用
    —— views                        //视图代码
        —— book
            —— home.html
        —— shared
            —— bookLayout.html      //书城模板页
—— main.go                          //主函数入口
```

book_tb表语句

```
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
```
