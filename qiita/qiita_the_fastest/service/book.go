package service

import (
	"qiita/model"
)

type BookService struct{}

func (BookService) SetBook(book *model.Book) {
	DbEngine.Create(book)
}

func (BookService) GetBookList() []model.Book {
	tests := make([]model.Book, 0)
	DbEngine.Distinct("id", "title", "content").Limit(10).Find(&tests)
	return tests
}

func (BookService) UpdateBook(newBook *model.Book) {
	DbEngine.Where("id = ?", newBook.Id).Updates(newBook)
}

func (BookService) DeleteBook(id int) {
	book := new(model.Book)
	DbEngine.Where("id = ?", id).Delete(book)
}
