package repository

import (
	"github.com/gtxiqbal/altera-agmc/Day-3/model"
)

type BookRepository interface {
	FindAll() []*model.Book
	FindByID(id uint) *model.Book
	Insert(book *model.Book)
	Update(book *model.Book)
	Delete(book *model.Book)
}
