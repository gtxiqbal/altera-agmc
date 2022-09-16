package repository

import (
	"github.com/gtxiqbal/altera-agmc/Day-4/model"
	"time"
)

type BookRepositoryImpl struct {
	books []*model.Book
}

func NewBookRepositoryImpl() *BookRepositoryImpl {
	bookRepo := &BookRepositoryImpl{}
	bookRepo.books = append(bookRepo.books, &model.Book{
		ID:        1,
		Title:     "Judul Buku Satu",
		Isbn:      "1-234-5678-9101112-13",
		Writer:    "Iqbal",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	})
	return bookRepo
}

func (bookRepo *BookRepositoryImpl) FindAll() []*model.Book {
	var books []*model.Book
	for _, book := range bookRepo.books {
		if book.DeletedAt == nil {
			books = append(books, book)
		}
	}
	return books
}

func (bookRepo *BookRepositoryImpl) FindByID(id uint) *model.Book {
	for _, book := range bookRepo.books {
		if book.ID == id && book.DeletedAt == nil {
			return book
		}
	}
	return nil
}

func (bookRepo *BookRepositoryImpl) Insert(book *model.Book) {
	var length uint = 0
	if len(bookRepo.books) > 0 {
		length = uint(len(bookRepo.books) - 1)
	}
	book.ID = bookRepo.books[length].ID + 1
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	book.DeletedAt = nil
	bookRepo.books = append(bookRepo.books, book)
}

func (bookRepo *BookRepositoryImpl) Update(bookUpdate *model.Book) {
	for _, book := range bookRepo.books {
		if book.ID == bookUpdate.ID && book.DeletedAt == nil {
			book.Title = bookUpdate.Title
			book.Isbn = bookUpdate.Isbn
			book.Writer = bookUpdate.Writer
			book.UpdatedAt = time.Now()
			return
		}
	}
}

func (bookRepo *BookRepositoryImpl) Delete(bookDelete *model.Book) {
	for _, book := range bookRepo.books {
		if book.ID == bookDelete.ID && book.DeletedAt == nil {
			now := time.Now()
			book.DeletedAt = &now
		}
	}
}
