package controllers

import (
	"errors"

	"demo.book.com/conf"
	"demo.book.com/models"
	"demo.book.com/services"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
)

type DemoController struct {
	Ctx iris.Context
	//Service services.SuperstarService
}

//自己实例化engine，获取单条数据
func (c *DemoController) GetRecord1() models.BookTb {
	engine, _ := xorm.NewEngine("mysql", "root:112233@tcp(127.0.0.1:3305)/mygo?charset=utf8")
	var info models.BookTb

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)

	engine.Table("book_tb").Where("id=?", 1).Get(&info)
	return info
}

//封装的单条记录
func (c *DemoController) GetOrm() iris.Map {
	//实例对象
	service := services.NewBookService()
	//ID获取单条数据
	info := service.Get(1)
	//获取列表
	list := service.GetList("Press = '湖南文艺出版社'", "ID asc", 2)
	//获取分页
	total, pageList := service.GetPageList("", "ID asc", 0, 2)
	//新增数据

	return iris.Map{
		"list":     list,
		"info":     info,
		"pageList": pageList,
		"total":    total,
	}
}

//返回xml
func (c *DemoController) GetXml() {
	service := services.NewBookService()
	//ID获取单条数据
	info := service.Get(1)
	c.Ctx.XML(info)
}

//故意报错
func (c *DemoController) GetErr() {
	//引发一个恐慌，程序会自动捕获并返回错误信息
	panic(errors.New("i'm a painc"))
}
func (c *DemoController) GetQps() string {
	return "hello"
}

func (c *DemoController) GetConf() map[string]string {
	reload := c.Ctx.URLParam("reload")
	if reload != "" {
		//如果有更新配置，重新读取配置文件
		conf.ReLoad()
	}

	return conf.SysConfMap
}
