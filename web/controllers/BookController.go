package controllers

import (
	"demo.book.com/conf"
	"demo.book.com/models"
	"demo.book.com/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type BookController struct {
	Ctx iris.Context
}

func (c *BookController) Get() mvc.Result {
	service := services.NewBookService()
	list := service.GetList("", "ID asc", 0)
	return mvc.View{
		Name: "book/home.html",
		Data: iris.Map{
			"Title":  "首页-" + conf.SysConfMap["port"],
			"List":   list,
			"Server": conf.SysConfMap["port"],
		},
		Layout: "shared/bookLayout.html",
	}
}

///book/ajaxbooks?key=go	访问地址是小写的
func (c *BookController) GetAjaxbooks() []models.BookTb {
	//获取url参数
	key := c.Ctx.URLParam("key")

	service := services.NewBookService()
	list := service.GetList(" bookName like '%"+key+"%'", "ID asc", 0)

	return list
}
