package dao

import (
	"demo.book.com/models"
	"github.com/go-xorm/xorm"
)

type BookDao struct {
	engine *xorm.Engine
}

//数据操作引擎
func NewBookDao(engine *xorm.Engine) *BookDao {
	return &BookDao{
		engine: engine,
	}
}

//获取单条
func (d *BookDao) Get(id int) *models.BookTb {
	data := &models.BookTb{Id: id}
	ok, err := d.engine.Table("book_tb").Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

//获取全部
//func (d *BookDao) GetAll() []models.BookTb {
//	datalist := make([]models.BookTb, 0)
//	err := d.engine.Table("book_tb").Desc("id").Limit(10, 0).Find(&datalist)
//	if err != nil {
//		return datalist
//	} else {
//		return datalist
//	}
//}

//获取列表
func (d *BookDao) GetList(query, sort string, pageSize int) []models.BookTb {
	datalist := make([]models.BookTb, 0)

	session := d.engine.Table("book_tb")
	if query != "" {
		session.Where(query)
	}
	if sort != "" {
		session.OrderBy(sort)
	}
	if pageSize > 0 {
		limit := pageSize
		start := 0
		session.Limit(limit, start)
	}
	err := session.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

//获取分页
func (d *BookDao) GetPageList(query, sort string, page, pageSize int) (int64, []models.BookTb) {
	datalist := make([]models.BookTb, 0)

	session := d.engine.Table("book_tb")
	if query != "" {
		session.Where(query)
	}
	if sort != "" {
		session.OrderBy(sort)
	}
	if pageSize > 0 {
		limit := pageSize
		start := page * pageSize
		session.Limit(limit, start)
	}
	total, err := session.FindAndCount(&datalist)

	if err != nil {
		return total, datalist
	} else {
		return total, datalist
	}
}

//删除
func (d *BookDao) Delete(id int) error {
	data := &models.BookTb{Id: id}
	_, err := d.engine.Table("book_tb").Delete(&data)
	return err
}

//更新
func (d *BookDao) Update(data *models.BookTb, columns []string) error {
	_, err := d.engine.Table("book_tb").ID(data.Id).MustCols(columns...).Update(data)
	return err
}

//新增
func (d *BookDao) Create(data *models.BookTb) error {
	_, err := d.engine.Table("book_tb").Insert(data)
	return err
}
