package services

import (
	"demo.book.com/dao"
	"demo.book.com/dbsource"
	"demo.book.com/models"
)

//书本服务接口
type IBookService interface {
	GetList(query, sort string, pageSize int) []models.BookTb
	GetPageList(query, sort string, page, pageSize int) (int64, []models.BookTb)
	Get(id int) *models.BookTb
	Delete(id int) error
	Update(user *models.BookTb, columns []string) error
	Create(user *models.BookTb) error
}

type bookService struct {
	dao *dao.BookDao
}

//创建数据操作服务
func NewBookService() IBookService {
	return &bookService{
		dao: dao.NewBookDao(dbsource.InstanceMaster()),
	}
}

func (s *bookService) GetList(query, sort string, pageSize int) []models.BookTb {
	return s.dao.GetList(query, sort, pageSize)
}

func (s *bookService) GetPageList(query, sort string, page, pageSize int) (int64, []models.BookTb) {
	return s.dao.GetPageList(query, sort, page, pageSize)
}

func (s *bookService) Get(id int) *models.BookTb {
	return s.dao.Get(id)
}
func (s *bookService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *bookService) Update(user *models.BookTb, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *bookService) Create(user *models.BookTb) error {
	return s.dao.Create(user)
}
